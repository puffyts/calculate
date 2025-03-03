package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/models"
	"github.com/solverANDimprover/calc_go/internal/repository"
	"github.com/solverANDimprover/calc_go/pkg/tools"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	e *echo.Echo
}

func NewHandler() *Handler {
	return &Handler{e: echo.New()}
}

func (h *Handler) AddExpression(ctx echo.Context) error {
	var req models.Request
	body, err := io.ReadAll(ctx.Request().Body)
	err = json.Unmarshal(body, &req)
	if ctx.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method is not allowed")
	}
	if err != nil || req.Expression == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid expression")
	}
	id := tools.NewCryptoRand()
	expr := models.NewExpression(id, false, 0, req)
	repository.Expressions.Add(id, expr)
	repository.NotInWork <- *expr
	ctx.Response().WriteHeader(http.StatusCreated)
	return nil
}
func (h *Handler) GetExpressions(ctx echo.Context) error {
	tasks := repository.Expressions.GetValues()
	ctx.Response().WriteHeader(http.StatusOK)
	return ctx.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetExpression(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty id value")
	}

	task, err := repository.Expressions.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Expression with id %d not found", id))
	}
	return ctx.JSONPretty(http.StatusOK, task, "\n")
}

func (h *Handler) UpdateExpression(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty id value")
	}
	var newReq models.Request
	err = json.NewDecoder(ctx.Request().Body).Decode(&newReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid expression")
	}
	repository.Expressions.Delete(id)
	repository.Expressions.Add(id, newReq)
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}

func (h *Handler) ExpressionForWork(ctx echo.Context) error {
	if len(repository.NotInWork) == 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "No available work")
	}
	expr := <-repository.NotInWork
	return ctx.JSON(http.StatusOK, expr)
}
