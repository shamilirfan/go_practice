package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=bookShop sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	log.Println("✅ Connected to Database successfully!")
}
