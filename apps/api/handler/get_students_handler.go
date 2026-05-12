package handler

import (
	"net/http"

	"go-demo/pkgs/models"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	studentID := r.URL.Query().Get("studentId")

	mu.Lock()
	defer mu.Unlock()

	if studentID != "" {
		s, ok := store[studentID]
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": "student not found",
			})
			return
		}
		writeJSON(w, http.StatusOK, s)
		return
	}

	students := make([]models.Student, 0, len(store))
	for _, s := range store {
		students = append(students, s)
	}
	writeJSON(w, http.StatusOK, students)
}
