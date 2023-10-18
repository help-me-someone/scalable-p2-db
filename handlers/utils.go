package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func PrepareHandler(w http.ResponseWriter, r *http.Request) (*json.Encoder, *gorm.DB, error) {
	response := json.NewEncoder(w)
	connection, ok := r.Context().Value("database").(*gorm.DB)
	if !ok {
		return nil, nil, fmt.Errorf("Could not connect to the database.")
	}
	return response, connection, nil
}
