package models

type Expression struct {
	ID      int     `json:"id"`
	Status  bool    `json:"status"`
	Result  float64 `json:"result"`
	Request `json:"request"`
}

type Request struct {
	Expression string `json:"expression"`
}

func NewExpression(id int, status bool, result float64, req Request) *Expression {
	return &Expression{ID: id, Status: status, Result: result, Request: req}
}
