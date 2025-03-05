package main

import "github.com/solverANDimprover/calc_go/internal/agent/service"

func main() {
	go service.Worker()
}
