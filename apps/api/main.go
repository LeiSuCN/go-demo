package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"go-demo/apps/api/handler"
)

//go:embed public/*
var staticFiles embed.FS

func main() {
	handler.InitStore()
	handler.AddTestData()

	subFS, err := fs.Sub(staticFiles, "public")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(subFS)))

	http.HandleFunc("/api/submit-student", handler.SubmitStudent)
	http.HandleFunc("/api/get-students", handler.GetStudents)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
