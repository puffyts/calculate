package calculation

import "errors"

var (
	ErrDivisionByZero        = errors.New("Division by zero")
	ErrUnmatchedOperator     = errors.New("Uncorrect type of operator")
	ErrMismatchedParentheses = errors.New("Missmatched Parentheses")
)
