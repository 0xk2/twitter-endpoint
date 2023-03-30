package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0xk2/twitter-endpoint/handler"
	"github.com/joho/godotenv"
)

var keys = make(map[string]string)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler.AuthHandler)
	log.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
