package main

import "github.com/solverANDimprover/calc_go/config"

func main() {
	application := config.NewApplication()
	application.StartServer()
}
