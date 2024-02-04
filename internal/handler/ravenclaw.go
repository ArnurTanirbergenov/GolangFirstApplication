package handler

import (
	"encoding/json"
    "net/http"
)

func RavenclawHandler(w http.ResponseWriter, r *http.Request) {
	house := House{"Когтевран", "Филиус Флитвик", "Роуэна Когтевран", "синий и бронзовый"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(house)
}
