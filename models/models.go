package models

import "time"

type Books struct {
	ID     int
	Name   string
	Author string
	Price  float32
	Image  string
}

type Customer struct {
	ID           int
	CustomerName string `json:"customer_name"`
	Email        string
}

type Order struct {
	ID            int         `json:"id"`
	CustomerID    int         `json:"customer_id"`
	RoadNumber    string      `json:"road_number"`
	HoldingNumber string      `json:"holding_number"`
	Area          string      `json:"area"`
	District      string      `json:"district"`
	PhoneNumber   string      `json:"phone_number"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	Items         []OrderItem `json:"items"`
}

type OrderItem struct {
	ID         int     `json:"id"`
	OrderID    int     `json:"order_id"`
	BookID     int     `json:"book_id"`
	BookName   string  `json:"book_name"`
	BookImage  string  `json:"book_image"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
