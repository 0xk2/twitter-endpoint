package main

import (
	"net/http"
	"os"

	"github.com/0xk2/twitter-endpoint/handler"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/verify", handler.VerifyHandler)
	http.ListenAndServe(":"+port, nil)
}
