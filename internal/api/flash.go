package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FlashRequest struct {
	Image  string `json:"image"`
	Device string `json:"device"`
	Config struct {
		SSH      bool   `json:"ssh"`
		Hostname string `json:"hostname"`
		Wifi     struct {
			SSID string `json:"ssid"`
			PSK  string `json:"psk"`
		} `json:"wifi"`
	} `json:"config"`
}

func FlashHandler(w http.ResponseWriter, r *http.Request) {
	var req FlashRequest
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("Flash request received:", req)

	// TODO: call internal/image + internal/device logic

	w.Write([]byte(`{"status":"ok"}`))
}
