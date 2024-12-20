package calculation_test

import (
	"github.com/solverANDimprover/calc_go/pkg"
	"testing"
)

type calculationTest struct {
	expression     string
	expectedresult float64
	err            error
}

func TestCalc(t *testing.T) {
	testCases := []calculationTest{
		{
			expression:     "2+2",
			expectedresult: 4,
			err:            nil,
		},
		{
			expression:     "22*4",
			expectedresult: 88,
			err:            nil,
		},
		{
			expression:     "180/16",
			expectedresult: 11.25,
			err:            nil,
		},
		{
			expression:     "(2+1) - (3-2) * (3+3)",
			expectedresult: -3,
			err:            nil,
		},
		{
			expression:     "((2+1)",
			expectedresult: 0,
			err:            calculation.ErrMismatchedParentheses,
		},
		{
			expression:     "7/0",
			expectedresult: 0,
			err:            calculation.ErrDivisionByZero,
		},
		{
			expression:     "+7+",
			expectedresult: 0,
			err:            calculation.ErrUnmatchedOperator,
		},
	}
	for _, test := range testCases {
		result, err := calculation.Calc(test.expression)
		if result != test.expectedresult {
			t.Fatalf("expected result: %.2f, got %.2f", test.expectedresult, result)
		}
		if err != test.err {
			t.Fatal("got", err, "expected:", test.err)
		}
	}
}
