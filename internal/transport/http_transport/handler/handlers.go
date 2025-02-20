package handler

import (
	"encoding/json"
	"fmt"
	"github.com/solverANDimprover/calc_go/internal/service"
	"github.com/solverANDimprover/calc_go/internal/transport/http_transport/models"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Service interface {
	Add(request models.Request) models.Expression
	GetExpressions() []models.Expression
}

type Handler struct {
	service  Service
	taskChan chan *models.Expression
	mutex    sync.Mutex
}

func NewHandler(service Service, ComputingPower int) *Handler {
	return &Handler{service: service, taskChan: make(chan *models.Expression, ComputingPower)}
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
	id := service.NewCryptoRand()
	expr := models.NewExpression(id, false, "in process", req)
	h.taskChan <- expr
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
	var tasks []models.Expression
	for taskPtr := range h.taskChan {
		tasks = append(tasks, *taskPtr)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/api/v1/expressions/"):])

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	for task := range h.taskChan {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}
