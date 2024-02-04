package app

import (
    "github.com/gorilla/mux"
	"github.com/ArnurTanirbergenov/GolangFirstApplication/internal/handler"
	"net/http"
)

func Start() {
    r := mux.NewRouter()

	r.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")
	
    r.HandleFunc("/gryffindor", handler.GryffindorHandler).Methods("GET")
    r.HandleFunc("/slytherin", handler.SlytherinHandler).Methods("GET")
    r.HandleFunc("/hufflepuff", handler.HufflepuffHandler).Methods("GET")
    r.HandleFunc("/ravenclaw", handler.RavenclawHandler).Methods("GET")

    http.ListenAndServe(":8080", r)
}