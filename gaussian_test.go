package gosig

import (
	"math"
	"testing"
)

func checkValue(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestGaussian(t *testing.T) {
	g := NewGaussian(1)
	coeffs := g.GetCoeffs(5)

	if !checkValue(coeffs[5], 0.3989, 0.0001) {
		t.Errorf("Coeffs[5] != 0.3989; got %f", coeffs[5])
	}
}
