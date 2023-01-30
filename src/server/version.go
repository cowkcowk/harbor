package server

import (
	"encoding/json"
	"net/http"
)

type APIVersion struct {
	Version string `json:"version"`
}

func GetAPIVersion(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&APIVersion{Version: ""}); err != nil {
		
	}
}