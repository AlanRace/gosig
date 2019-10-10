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
	dictionary := NewGaussianDictionary(gaussian, len(data))

	fmt.Println(dictionary)
	fmt.Println(dictionary.middleEntry)

	rows, cols := dictionary.Dims()

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Printf("%.7f   ", dictionary.At(r, c))
		}

		fmt.Println()
	}

	fmt.Println(OMP(&signal, dictionary, 5, 3))
}
