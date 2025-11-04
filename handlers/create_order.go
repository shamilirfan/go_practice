package handlers

import (
	"database/sql"
	"fmt"
	"go_practice/models"
)

func CreateOrder(db *sql.DB, order *models.Order) (err error) {
	query := `
		INSERT INTO orders (
			customer_id, 
			road_number,
			holding_number,
			area, district, 
			phone_number,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6,'Pending')
		RETURNING id, created_at
	`
	/*
		👉 QueryRow() এর কাজ হলো — একটা single row (একটা record) ডাটাবেস থেকে নিয়ে আসা।
			একটি row রিটার্ন করে। যদি query তে একাধিক row ফেরত আসে, শুধু প্রথমটা নেয়।

		👉 Scan() এর কাজ হলো — QueryRow() থেকে পাওয়া সেই একটি row-এর কলাম ভ্যালুগুলোকে
			Go ভ্যারিয়েবলগুলোর মধ্যে কপি করে দেওয়া।
	*/
	err = db.QueryRow(query,
		order.CustomerID,
		order.RoadNumber,
		order.HoldingNumber,
		order.Area,
		order.District,
		order.PhoneNumber,
	).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		return err
	}

	// total
	var total float64

	for i := range order.Items {
		var price float64
		var name, image string
		err = db.QueryRow(`SELECT name, price, image FROM books WHERE id = $1`,
			order.Items[i].BookID,
		).Scan(&name, &price, &image)
		if err != nil {
			return fmt.Errorf("book %d not found: %w", order.Items[i].BookID, err)
		}

		order.Items[i].BookName = name
		order.Items[i].BookImage = image
		order.Items[i].UnitPrice = price
		order.Items[i].TotalPrice = price * float64(order.Items[i].Quantity)
		total += order.Items[i].TotalPrice

		query = `
			INSERT INTO order_items (order_id, book_id, book_name, book_image, quantity, unit_price, total_price)
			VALUES ($1, $2, $3, $4, $5,  $6,  $7)
		`
		/*
			👉 db.Exec() কোনো row ফেরত দেয় না — শুধু ডাটাবেসে একটা রেকর্ড insert, update, delete করে।
				db.Exec() দুইটা জিনিস ফেরত দেয়:
				1.
					. sql.Result → এতে থাকে: LastInsertId() — সর্বশেষ insert করা row-এর ID (যদি DB সমর্থন করে)
					. RowsAffected() — কতগুলো row পরিবর্তন হয়েছে।
				2.	.err → error হলে সেটা ধরে রাখে।
		*/
		_, err = db.Exec(
			query,
			order.ID,
			order.Items[i].BookID,
			order.Items[i].BookName,
			order.Items[i].BookImage,
			order.Items[i].Quantity,
			order.Items[i].UnitPrice,
			order.Items[i].TotalPrice,
		)
		if err != nil {
			return fmt.Errorf("failed to insert order item: %w", err)
		}
	}

	order.Status = "Pending"
	return nil
}
