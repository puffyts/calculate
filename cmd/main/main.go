package main

import (
	"github.com/solverANDimprover/calc_go/internal/config"
)

func main() {
	application := config.NewApplication()
	application.StartServer()
}
