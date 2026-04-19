package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	//api endpoints
	mux.HandleFunc("/api/devices", DevicesHandler)
	mux.HandleFunc("/api/flash", FlashHandler)
	mux.HandleFunc("/api/status", StatusHandler)

	return mux
}
