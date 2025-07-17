package main

import (
	"fmt"
	"net/http"

	"crud-web-server/db"
	"crud-web-server/handlers"
)

func main() {
	db.InitDB()

	http.HandleFunc("/user/insert", handlers.InsertUserHandler)
	http.HandleFunc("/user/update", handlers.UpdateUserHandler)

	fmt.Println("Server running on http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}
