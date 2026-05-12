package handler

import (
	"encoding/json"
	"net/http"

	"go-demo/pkgs/models"
)

func SubmitStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if s.ID == "" || s.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "id and name are required",
		})
		return
	}

	mu.Lock()
	store[s.ID] = s
	mu.Unlock()

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "student saved successfully",
	})
}
