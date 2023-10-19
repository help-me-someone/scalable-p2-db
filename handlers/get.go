package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
			"message": "User not found.",
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

func GetUserVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
			"message": "User not found.",
		})
		return
	}

	vids, err := crud.GetUserVideos(connection, usr.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get videos.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "Found videos.",
		"videos":  vids,
	})
}

func GetUserVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	videoName := p.ByName("video")

	usr, err := crud.GetUserByName(connection, username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "User not found.",
		})
		return
	}

	vid, err := crud.GetUserVideo(connection, videoName, usr.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get videos.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "Found videos.",
		"video":   vid,
	})
}

func GetVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, connection, err := PrepareHandler(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Could not connect to the database.",
		})
		return
	}

	key := p.ByName("key")
	vid, err := crud.GetVideoByKey(connection, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Video not found.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "Video found.",
		"video":   vid,
	})
}

func GetTopPopularVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, connection, err := PrepareHandler(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Could not connect to the database.",
		})
		return
	}

	// Attemp to get the query values.
	amountStr := p.ByName("amount")
	pageStr := p.ByName("page")
	amountEmpty := len(amountStr) == 0
	pageEmpty := len(pageStr) == 0
	if amountEmpty || pageEmpty {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get videos.",
		})
		return
	}

	// Convert query into numerical values.
	amount, err := strconv.ParseInt(amountStr, 10, 32)
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get videos.",
		})
	}

	vids, err := crud.GetTopPopularVideos(connection, int(page), int(amount))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get videos.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Encode(map[string]interface{}{
		"success": true,
		"message": "Found videos.",
		"videos":  vids,
	})
}
