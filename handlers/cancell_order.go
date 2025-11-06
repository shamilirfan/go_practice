package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CancellOrder(r *http.Request, db *sql.DB) (err error) {
	str := r.PathValue("id")
	id, _ := strconv.Atoi(str)

	var status struct {
		Status string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		fmt.Println(err)
		return
	}

	query := `
	UPDATE orders
	SET status = $1
	WHERE id = $2
	`
	_, err = db.Exec(query, status.Status, id)
	if err != nil {
		return err
	}

	return nil
}
