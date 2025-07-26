package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	customers, err := getAllCustomers()
	if err != nil {
		http.Error(w, "Error fetching customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetCustomerHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	customerIdStr := r.URL.Query().Get("customerid")
	if customerIdStr == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}
	customer, err := getCustomer(customerId)
	if err != nil {
		http.Error(w, "Error fetching customer: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if customer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload []map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload) == 0 {
		http.Error(w, "No data provided for update", http.StatusBadRequest)
		return
	}
	if err := updateCustomers(payload); err != nil {
		http.Error(w, "Failed to update customers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Customers update successful"))
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	customerIdStr := r.URL.Query().Get("customerid")
	if customerIdStr == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}

	if err := deleteCustomer(customerId); err != nil {
		http.Error(w, "Failed to delete customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Customer is been deleted successfully"))
}
