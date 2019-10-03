package gosig

import (
	"gonum.org/v1/gonum/mat"
)

// ConvolveMatVec performs convolution on the data matrix, convolving the vector with each row in the matrix (moving across each row).
// A new matrix is created and returned with the result, the input matrices are unaltered.
func ConvolveMatVec(data *mat.Dense, vec *mat.VecDense) *mat.Dense {
	rows, cols := data.Dims()
	result := mat.NewDense(rows, cols, nil)

	vecLen := vec.Len()
	halfWidth := (vecLen - 1) / 2

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for x := -halfWidth; x <= halfWidth; x++ {
				if x+col < 0 {
					continue
				} else if x+col >= cols {
					break
				}

				result.Set(row, col, result.At(row, col)+data.At(row, x+col)*vec.AtVec(x+halfWidth))
			}
		}
	}

	return result
}
