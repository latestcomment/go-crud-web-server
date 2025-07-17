package handlers

import (
	"crud-web-server/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"crud-web-server/models"
)

func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post is allowed", http.StatusMethodNotAllowed)
	}
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	if err := InsertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := map[string]interface{}{
		"message": "User inserted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT is allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request "+err.Error(), http.StatusBadRequest)
	}

	idRaw, ok := payload["id"]
	if !ok {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idRaw.(string))
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	delete(payload, "ID")

	if err := UpdateUser(id, payload); err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
	}

	resp := map[string]interface{}{
		"message": "User updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func InsertUser(user models.User) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func UpdateUser(id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setClauses := []string{}
	values := []interface{}{}

	i := 1
	for col, val := range updates {
		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, i))
		values = append(values, val)
		i++
	}

	values = append(values, id)
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(setClauses, ", "), i)

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, values...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
