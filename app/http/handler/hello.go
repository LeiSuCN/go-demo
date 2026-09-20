package handler

import (
	"fmt"
	"net/http"

	"go-demo/model"
)

// Hello responds with a plain-text greeting.
func Hello(w http.ResponseWriter, r *http.Request) {
	msg := model.Message{Text: "Hello, world!"}
	fmt.Fprintln(w, msg.Text)
}
