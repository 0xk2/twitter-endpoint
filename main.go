package main

import (
	"net/http"

	"github.com/0xk2/twitter-endpoint/handler"
)

func main() {
	http.HandleFunc("/verify", handler.VerifyHandler)
	http.ListenAndServe(":8080", nil)
}
