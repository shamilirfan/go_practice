package handlers

import (
	"database/sql"
	"fmt"
	"go_practice/models"
)

func CreateOrder(db *sql.DB, order *models.Order) (err error) {
	query := `
		INSERT INTO orders (
			customer_name, 
			road_number,
			holding_number,
			area, district, 
			phone_number,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6,'Pending')
		RETURNING id, created_at
	`
	err = db.QueryRow(query,
		order.CustomerName,
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
		order.Items[i].Subtotal = price * float64(order.Items[i].Quantity)
		total += order.Items[i].Subtotal

		query = `
			INSERT INTO order_items (order_id, book_id, book_name, book_image, quantity, unit_price, subtotal)
			VALUES ($1, $2, $3, $4, $5,  $6,  $7)
		`
		_, err = db.Exec(
			query,
			order.ID,
			order.Items[i].BookID,
			order.Items[i].BookName,
			order.Items[i].BookImage,
			order.Items[i].Quantity,
			order.Items[i].UnitPrice,
			order.Items[i].Subtotal,
		)
		if err != nil {
			return fmt.Errorf("failed to insert order item: %w", err)
		}
	}

	query = `UPDATE orders SET total_price = $1 WHERE id = $2`
	_, err = db.Exec(query, total, order.ID)
	if err != nil {
		return err
	}

	order.TotalPrice = total
	order.Status = "Pending"
	return nil
}
