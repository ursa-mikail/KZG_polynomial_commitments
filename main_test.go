package main

import (
	"math/big"
	"testing"
)

func TestEvaluatePolynomial(t *testing.T) {
	p := []*big.Int{
		big.NewInt(5),  // x^0
		big.NewInt(1),  // x^1
		big.NewInt(0),  // x^2
		big.NewInt(1),  // x^3
	}
	z := 3
	expected := 5 + 1*3 + 0*9 + 1*27 // = 35

	if got := evaluatePolynomial(p, z); got != expected {
		t.Errorf("evaluatePolynomial() = %v, want %v", got, expected)
	}
}

