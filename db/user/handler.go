package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post is allowed", http.StatusMethodNotAllowed)
	}
	var u User
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
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created")
}
