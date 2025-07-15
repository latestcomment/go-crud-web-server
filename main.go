package main

import (
	"fmt"
	"net/http"

	"crud-web-server/db"
	"crud-web-server/db/user"
)

func main() {
	db.InitDB()

	http.HandleFunc("/user/insert", user.InsertUserHandler)

	fmt.Println("Server running on http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}
