package gosig

import (
	"fmt"
	"testing"
)

func TestOMP(t *testing.T) {
	var data []float64
	data = make([]float64, 10)

	for i := 0; i < len(data); i++ {
		data[i] = float64(i + 1)
	}

	signal := Signal{X: data, Y: data}

	gaussian := NewGaussian(1.0)
	dictionary := NewGaussianDictionary(gaussian)

	fmt.Println(OMP(&signal, dictionary, 5, 3))
}
