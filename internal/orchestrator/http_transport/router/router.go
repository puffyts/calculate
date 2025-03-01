package router

import (
	"github.com/labstack/echo/v4"
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/handler"
	"log"
	"net/http"
)

type Config struct {
	Port string
}

type Router struct {
	Router  *echo.Echo
	Handler handler.Handler
	config  Config
}

func NewRouter(config Config, h *handler.Handler) *Router {
	e := echo.New()
	e.POST("/api/v1/calculate", h.AddExpression)
	e.GET("/api/v1/expressions", h.GetExpressions)
	e.GET("/api/v1/expressions/:id", h.GetExpression)
	e.POST("/internal/task", h.UpdateExpression)
	e.GET("/internal/task", h.ExpressionForWork)
	return &Router{config: config, Router: e}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(":"+r.config.Port, r.Router))
}
