package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type face struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var faces []face

func main() {
	http.HandleFunc("/faces", handleGetFaces)
	http.HandleFunc("/add", handleAddFaces)
	http.HandleFunc("/update", handleUpdateFunc)
	http.HandleFunc("/delete", handleDeleteFunc)

	fmt.Println("Server running at http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}

func handleGetFaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, faces)
}

func handleAddFaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post is allowed", http.StatusMethodNotAllowed)
		return
	}
	var newFaces []face
	if err := decodeJSON(r.Body, &newFaces); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := range newFaces {
		if newFaces[i].ID == uuid.Nil {
			newFaces[i].ID = uuid.New()
		}
	}
	faces = addFaces(newFaces, faces)
	writeJSON(w, faces)

}

func handleUpdateFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only Put method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var updates []face
	if err := decodeJSON(r.Body, &updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	faces = updateFaces(updates, faces)
	writeJSON(w, faces)
}

func handleDeleteFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only Delete method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var ids []uuid.UUID
	if err := decodeJSON(r.Body, &ids); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	faces = deleteFaces(ids, faces)
	writeJSON(w, faces)
}

// Utility function
func decodeJSON(body io.ReadCloser, target interface{}) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(target)
}
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Logic function
func addFaces(NewFaces []face, faces []face) []face {
	faces = append(faces, NewFaces...)

	return faces
}

func deleteFaces(ids []uuid.UUID, faces []face) []face {
	idSet := make(map[uuid.UUID]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	newFaces := []face{}

	for _, f := range faces {
		if _, isFound := idSet[f.ID]; !isFound {
			newFaces = append(newFaces, f)
		}
	}
	return newFaces
}

func updateFaces(updates []face, faces []face) []face {
	updateMap := make(map[uuid.UUID]string)
	for _, f := range updates {
		updateMap[f.ID] = f.Name
	}

	for i, f := range faces {
		if newName, found := updateMap[f.ID]; found {
			faces[i].Name = newName
		}
	}
	return faces
}
