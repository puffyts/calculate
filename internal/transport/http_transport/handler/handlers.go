package handler

import (
	"encoding/json"
	"fmt"
	"github.com/solverANDimprover/calc_go/internal/repository"
	"github.com/solverANDimprover/calc_go/internal/transport/http_transport/models"
	"github.com/solverANDimprover/calc_go/pkg/tools"
	"io"
	"net/http"
	"strconv"
)

type Service interface {
	Add(request models.Request) models.Expression
	GetExpressions() []models.Expression
}

type Handler struct {
	service Service
}

func NewHandler(service Service, ComputingPower int) *Handler {
	return &Handler{service: service}
}

func (h *Handler) addTask(w http.ResponseWriter, r *http.Request) {
	var req models.Request
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
	id := tools.NewCryptoRand()
	expr := models.NewExpression(id, false, "in process", req)
	repository.Tasks.Add(id, expr)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, fmt.Sprintf(`{"id":%d}`, id))

	//var response = new(models.Expression)
	//response.Result = res
	//responseJson, err := json.Marshal(response)
	//if err != nil {
	//	http.Error(w, `{error: "Internal server error"}`, 500)
	//	return
	//}
	//w.Write(responseJson)
}
func (h *Handler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := repository.Tasks.GetValues()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/api/v1/expressions/"):])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	task, err := repository.Tasks.Get(id)
	if err != nil {
		http.Error(w, "Taкой задачи нет", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}
