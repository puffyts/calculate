package config

import (
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/router"
)

type Config struct {
	RouterConfig router.Config
}

func New() *Config {
	var cfg Config
	cfg.RouterConfig = router.Config{Port: "8080"}
	return &cfg
}
