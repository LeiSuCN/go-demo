package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"go-demo/pkgs/models"
)

var (
	store = make(map[string]models.Student)
	mu    sync.Mutex
)

func InitStore() {
	mu.Lock()
	defer mu.Unlock()
	store = make(map[string]models.Student)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
