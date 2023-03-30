package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Success bool `json:"success"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	data := ResponseData{
		Success: true,
	}
	jsondata, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsondata)
}
