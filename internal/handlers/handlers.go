package handlers

import (
	"encoding/json"
	request2 "github.com/solverANDimprover/calc_go/internal/request"
	"github.com/solverANDimprover/calc_go/pkg/calculation"
	"io"
	"net/http"
)

type request struct {
	Expression string `json:"expression"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req request
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &req)
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method is not allowed"}`, 405)
		return
	}
	if err != nil || req.Expression == "" {
		http.Error(w, `{"error":"Expression is not valid"}`, 422)
		return
	}
	res, err := calculation.Calc(req.Expression)
	var response *request2.Response = new(request2.Response)
	response.Result = res
	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{error: "Internal server error"")`, 500)
	}
	w.Write(responseJson)
}
