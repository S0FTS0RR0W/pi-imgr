package main

import (
	"log"
	"net/http"

	"charlix.dev/pi-imgr/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("pi-imgr backend running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
