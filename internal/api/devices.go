package api

import (
	"encoding/json"
	"net/http"
)

type Device struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Size string `json:"size"`
}

// to be replaced with actual logic after frontend is ready

func DevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices := []Device{
		{Path: "dev/sda", Name: "Disk 1", Size: "500GB"},
		{Path: "/dev/sdb", Name: "Disk 2", Size: "1TB"},
	}

	json.NewEncoder(w).Encode(devices)
}
