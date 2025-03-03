package service

import (
	"github.com/solverANDimprover/calc_go/internal/agent/client"
	repo "github.com/solverANDimprover/calc_go/internal/repository"
	"os"
	"strconv"
	"time"
)

func Work() {
	for {
		req, err := client.GetWork()
		if err != nil {
			time.Sleep(time.Second * 5)
		}
		res, err := Calc(req.Request.Expression)
		req.Result = res
		if err != nil {
			repo.InvalidExprChan <- req
		}
		repo.DoneExprChan <- req
		time.Sleep(time.Second * 1)
	}
}

func Worker() {
	CompPower, _ := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	for i := 0; i < CompPower; i++ {
		go Work()
	}
}
