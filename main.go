package main

import (
	"crud-web-server/db"
	"crud-web-server/handlers"
	"fmt"
	"net/http"
)

func main() {
	db.InitDB()

	// http.HandleFunc("/user/insert", handlers.InsertUserHandler)
	// http.HandleFunc("/user/update", handlers.UpdateUserHandler)

	http.HandleFunc("/customers", handlers.GetAllCustomersHandler)
	http.HandleFunc("/customer", handlers.GetCustomerHanlder)
	http.HandleFunc("/customer/update", handlers.UpdateCustomerHandler)
	http.HandleFunc("/customer/delete", handlers.DeleteCustomerHandler)
	fmt.Println("Server running on http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}
