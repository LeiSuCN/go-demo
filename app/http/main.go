package main

import (
	"log"
	"net/http"

	"go-demo/app/http/handler"
)

func main() {
	http.HandleFunc("/", handler.Hello)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
