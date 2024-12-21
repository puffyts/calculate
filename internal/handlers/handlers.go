package handlers

import (
	"encoding/json"
	"github.com/solverANDimprover/calc_go/internal/funcs"
	"net/http"
)

type request struct {
	expression string `json:"expression"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.expression == "" {
		w.WriteHeader(http.StatusBadRequest)
		funcs.SendErrJson(w)
	}
}
