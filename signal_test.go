package gosig

import (
	"fmt"
	"testing"
)

func TestGradient(t *testing.T) {
	data := []float64{0.4505, 0.0838, 0.2290, 0.9133, 0.1524, 0.8258, 0.5383, 0.9961, 0.0782, 0.4427}
	result := []float64{-0.3667, -0.1108, 0.4148, -0.0383, -0.0438, 0.1930, 0.0852, -0.2301, -0.2767, 0.3645}

	fmt.Println(data)
	fmt.Println(Gradient(data))
	fmt.Println(result)
}
