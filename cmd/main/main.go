package main

import (
	"github.com/solverANDimprover/calc_go/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	http.ListenAndServe(":8000", mux)
}
