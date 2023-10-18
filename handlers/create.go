package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/help-me-someone/scalable-p2-db/functions/crud"
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, connection, err := PrepareHandler(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Could not connect to the database.",
		})
		return
	}

	// Retrieve the user information.
	usr := &user.User{}
	err = json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Could not decode request.",
		})
		return
	}

	// Username/Password not given.
	if len(usr.Username) == 0 || len(usr.HashedPassword) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Username/Password not specified.",
		})
		return
	}

	// Create the entry.
	_, err = crud.CreateUser(connection, usr.Username, usr.HashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Could not create user.",
		})
		return
	}

	// Everything went well.
	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "User successfully created.",
	})
}
