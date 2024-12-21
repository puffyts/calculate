package handlers

import (
	"encoding/json"
	request2 "github.com/solverANDimprover/calc_go/internal/request"
	"github.com/solverANDimprover/calc_go/pkg/calculation"
	"net/http"
)

type request struct {
	expression string `json:"expression"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method is not allowed"}`, 405)
	}
	if err != nil || req.expression == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, `{"error":"Expression is not valid"}`, 422)
	}
	res, err := calculation.Calc(req.expression)
	var response *request2.Response = new(request2.Response)
	response.Result = res
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(responseJson)
}
