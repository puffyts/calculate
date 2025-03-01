package repository

import "github.com/solverANDimprover/calc_go/pkg/SyncMap"

var Expressions = SyncMap.NewSyncMap()
var InWork []int
var NotInWork []int
