package calculation

import (
	"fmt"
	"github.com/solverANDimprover/calc_go/pkg/errors"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	var operators []rune
	var operands []float64

	priority := func(op rune) int {
		if op == '*' || op == '/' {
			return 2
		} else if op == '+' || op == '-' {
			return 1
		}
		return 0
	}

	calculate := func() error {
		if len(operators) == 0 {
			return errors.ErrInvalidExpression
		}

		op := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if len(operands) < 2 {
			return errors.ErrInvalidExpression
		}

		b := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		a := operands[len(operands)-1]
		operands = operands[:len(operands)-1]

		var result float64
		switch op {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return errors.ErrDivisionByZero
			}
			result = a / b
		default:
			return errors.ErrInvalidExpression
		}

		operands = append(operands, result)
		return nil
	}

	var numStr string
	for _, char := range expression {
		if unicode.IsSpace(char) {
			continue
		}
		if unicode.IsDigit(char) || char == '.' {
			numStr += string(char)
		} else {
			if numStr != "" {
				num, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return 0, fmt.Errorf("invalid number: %s", numStr)
				}
				operands = append(operands, num)
				numStr = ""
			}

			if char == '+' || char == '-' || char == '*' || char == '/' {
				for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(char) {
					if err := calculate(); err != nil {
						return 0, err
					}
				}
				operators = append(operators, char)
			} else if char == '(' {
				operators = append(operators, char)
			} else if char == ')' {
				for len(operators) > 0 && operators[len(operators)-1] != '(' {
					if err := calculate(); err != nil {
						return 0, err
					}
				}
				if len(operators) == 0 || operators[len(operators)-1] != '(' {
					return 0, errors.ErrMismatchedParentheses
				}
				operators = operators[:len(operators)-1]
			} else {
				return 0, fmt.Errorf("invalid character: %c", char)
			}
		}
	}

	if numStr != "" {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", numStr)
		}
		operands = append(operands, num)
	}

	for len(operators) > 0 {
		if err := calculate(); err != nil {
			return 0, err
		}
	}

	if len(operands) != 1 {
		return 0, errors.ErrInvalidExpression
	}

	return operands[0], nil
}
