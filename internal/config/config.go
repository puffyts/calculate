package config

import (
	"github.com/solverANDimprover/calc_go/internal/transport/http_transport/handler"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

type Application struct {
	config      *Config
	INFOLogger  *log.Logger
	ERRORLogger *log.Logger
}

func NewConfig() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8000"
	}
	return config
}

func NewApplication() *Application {
	return &Application{config: NewConfig(),
		INFOLogger:  log.New(os.Stdout, "[INFO]", 4000),
		ERRORLogger: log.New(os.Stdout, "[ERROR]", 4000)}
}

func (a *Application) StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
	a.INFOLogger.Printf("Server started at :%s", a.config.Addr)
	a.ERRORLogger.Fatal(http.ListenAndServe(":"+a.config.Addr, mux))
}
