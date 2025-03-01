package models

type Expression struct {
	ID      int    `json:"id"`
	Status  bool   `json:"status"`
	Result  string `json:"result"`
	Request `json:"-"`
}

type Request struct {
	Expression string `json:"expression"`
}

func NewExpression(id int, status bool, result string, req Request) *Expression {
	return &Expression{ID: id, Status: status, Result: result, Request: req}
}
