package repository

import (
	"github.com/solverANDimprover/calc_go/internal/orchestrator/http_transport/models"
	"github.com/solverANDimprover/calc_go/pkg/SyncMap"
)

var Expressions = SyncMap.NewSyncMap()
var InWork []int
var NotInWork chan models.Expression
var DoneExprChan chan models.Expression
var InvalidExprChan chan models.Expression
