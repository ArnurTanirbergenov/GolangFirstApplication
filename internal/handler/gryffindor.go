package handler

import (
	"encoding/json"
    "net/http"
)

func GryffindorHandler(w http.ResponseWriter, r *http.Request) {
	house := House{"Гриффиндор", "Минерва МакГонагалл", "Годрик Гриффиндор", "красный и золотой"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(house)
}