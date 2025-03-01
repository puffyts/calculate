package service

import "errors"

var (
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrDivisionByZero        = errors.New("division by zero")
	ErrEmptyInput            = errors.New("empty input")
	ErrMismatchedParentheses = errors.New("mismatched parentheses")
	ErrInvalidNumber         = errors.New("invalid number")
	ErrUnexpectedToken       = errors.New("unexpected token")
	ErrNotEnoughValues       = errors.New("not enough values in expression")
	ErrInvalidOperator       = errors.New("invalid operator")
	ErrOperatorAtEnd         = errors.New("operator at end of expression")
	ErrMultipleDecimalPoints = errors.New("multiple decimal points in a number")
)
