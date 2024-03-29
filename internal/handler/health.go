package handler

import (
	"encoding/json"
    "net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request){
	health := struct{
		Status string `json:status`
	}{
		Status:"OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(health)
}
