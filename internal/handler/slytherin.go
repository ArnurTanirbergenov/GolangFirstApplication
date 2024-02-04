package handler

import (
	"encoding/json"
    "net/http"
)

func SlytherinHandler(w http.ResponseWriter, r *http.Request) {
	house := House{"Слизерин", "Северус Снегг", "Салазар Слизерин", "зелёный и серебряный"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(house)
}