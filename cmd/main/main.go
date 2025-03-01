package main

import (
	"github.com/solverANDimprover/calc_go/internal/config"
	handler2 "github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/handler"
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/router"
)

func main() {
	cfg := config.New()
	handler := handler2.NewHandler()
	router := router.NewRouter(cfg.RouterConfig, handler)
	router.Run()

}
