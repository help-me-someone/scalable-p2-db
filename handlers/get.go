package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/help-me-someone/scalable-p2-db/functions/crud"
	"github.com/julienschmidt/httprouter"
)

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, connection, err := PrepareHandler(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Could not connect to the database.",
		})
		return
	}

	username := p.ByName("username")

	usr, err := crud.GetUserByName(connection, username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Could not find user by username.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "User found.",
		"user":    usr,
	})
}
