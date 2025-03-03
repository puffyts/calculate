package client

import (
	"encoding/json"
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/models"
	"net/http"
)

func GetWork() (models.Expression, error) {
	resp, err := http.Get("localhost/internal/task")
	if err != nil {
		return models.Expression{}, err
	}
	var expr models.Expression
	err = json.NewDecoder(resp.Body).Decode(&expr)
	if err != nil {
		return models.Expression{}, err
	}
	return expr, err
}
