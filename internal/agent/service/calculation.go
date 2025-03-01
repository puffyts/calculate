package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	addition, _    = time.ParseDuration(os.Getenv("TIME_ADDITION_MS"))
	subtraction, _ = time.ParseDuration(os.Getenv("TIME_ADDITION_MS"))
	multiplier, _  = time.ParseDuration(os.Getenv("TIME_ADDITION_MS"))
	divisor, _     = time.ParseDuration(os.Getenv("TIME_ADDITION_MS"))
)

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, ErrEmptyInput
	}
	expression = strings.ReplaceAll(expression, " ", "")
	return evaluateExpression(expression)
}

func evaluateExpression(expression string) (float64, error) {
	var values []float64
	var ops []rune
	i := 0
	for i < len(expression) {
		char := rune(expression[i])
		if char == '(' {
			ops = append(ops, char)
			i++
			continue
		} else if char == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOperation(&values, &ops); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, ErrMismatchedParentheses
			}
			ops = ops[:len(ops)-1]
			i++
			continue
		} else if isDigit(char) ||
			(char == '-' && (i == 0 || expression[i-1] == '(' || isOperator(rune(expression[i-1])))) ||
			(char == '+' && i == 0) {
			start := i
			if char == '-' || char == '+' {
				i++
			}
			for i < len(expression) && (isDigit(rune(expression[i])) || expression[i] == '.') {
				if expression[i] == '.' && (i < len(expression)-1 && expression[i+1] == '.') {
					return 0, ErrMultipleDecimalPoints
				}
				i++
			}
			num, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s: %w", expression[start:i], ErrInvalidNumber)
			}
			values = append(values, num)
			continue
		} else if isOperator(char) {
			if i == len(expression)-1 {
				return 0, ErrOperatorAtEnd
			}
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(char) {
				if err := applyOperation(&values, &ops); err != nil {
					return 0, err
				}
			}
			ops = append(ops, char)
			i++
			continue
		} else {
			return 0, fmt.Errorf("unexpected token: %s: %w", string(char), ErrUnexpectedToken)
		}
	}

	for len(ops) > 0 {
		if err := applyOperation(&values, &ops); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, ErrInvalidExpression
	}
	return values[0], nil
}

func applyOperation(values *[]float64, ops *[]rune) error {
	if len(*values) < 2 || len(*ops) == 0 {
		return ErrNotEnoughValues
	}
	b := (*values)[len(*values)-1]
	a := (*values)[len(*values)-2]
	op := (*ops)[len(*ops)-1]

	*values = (*values)[:len(*values)-2]
	*ops = (*ops)[:len(*ops)-1]

	switch op {
	case '+':
		*values = append(*values, a+b)
		time.Sleep(time.Second * addition)
	case '-':
		*values = append(*values, a-b)
		time.Sleep(time.Second * subtraction)
	case '*':
		*values = append(*values, a*b)
		time.Sleep(time.Second * multiplier)
	case '/':
		if b == 0 {
			return ErrDivisionByZero
		}
		*values = append(*values, a/b)
		time.Sleep(time.Second * divisor)
	default:
		return fmt.Errorf("invalid operator: %v: %w", string(op), ErrInvalidOperator)
	}

	return nil
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}
