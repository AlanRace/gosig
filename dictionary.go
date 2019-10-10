package gosig

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Dictionary interface {
	mat.Matrix
}

type PartialDictionary struct {
	dict Dictionary

	cols []int
	rows []int
}

func NewPartialDictionary(dict Dictionary) *PartialDictionary {
	var partDict PartialDictionary
	partDict.dict = dict

	rows, cols := dict.Dims()

	partDict.rows = make([]int, rows)
	for r := 0; r < rows; r++ {
		partDict.rows[r] = r
	}

	partDict.cols = make([]int, cols)
	for c := 0; c < cols; c++ {
		partDict.cols[c] = c
	}

	return &partDict
}

// Dims returns the dimensions of a Matrix.
func (dict *PartialDictionary) Dims() (r, c int) {
	return len(dict.rows), len(dict.cols)
}

func (dict *PartialDictionary) At(i, j int) float64 {
	return dict.dict.At(dict.rows[i], dict.cols[j])
}

func (dict *PartialDictionary) T() mat.Matrix {
	return &PartialDictionary{dict: dict.dict.T(), rows: dict.cols, cols: dict.rows}
}

func (dict *PartialDictionary) KeepCols(cols []int) {
	dict.cols = cols
}

type GaussianDictionary struct {
	numEntries int
	gaussian   *Gaussian

	middleEntry *mat.VecDense
	halfWidth   int
}

// Dims returns the dimensions of a Matrix.
func (dict *GaussianDictionary) Dims() (r, c int) {
	return dict.numEntries, dict.numEntries
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// At returns the value of a matrix element at row i, column j.
// It will panic if i or j are out of bounds for the matrix.
func (dict *GaussianDictionary) At(i, j int) float64 {
	left := min(i, j)
	right := max(i, j)

	index := (left - right) + dict.halfWidth

	if index >= dict.middleEntry.Len() || index < 0 {
		return 0.0
	}

	return dict.middleEntry.AtVec(index)
}

// T returns the transpose of the Matrix. Whether T returns a copy of the
// underlying data is implementation dependent.
// This method may be implemented using the Transpose type, which
// provides an implicit matrix transpose.
func (dict *GaussianDictionary) T() mat.Matrix {
	return dict
}

func NewGaussianDictionary(gaussian *Gaussian, numEntries int) *GaussianDictionary {
	var dict GaussianDictionary
	dict.numEntries = numEntries
	dict.gaussian = gaussian

	dict.halfWidth = ((numEntries - 1) / 2) + 1

	coeffVals := dict.gaussian.GetCoeffs(int(math.Ceil(dict.gaussian.sigma) * 5))
	coeffs := mat.NewVecDense(len(coeffVals), coeffVals)

	dict.middleEntry = mat.NewVecDense(dict.halfWidth*2+1, nil)
	dict.middleEntry.SetVec(dict.halfWidth, 1)

	dict.middleEntry = ConvolveVecVec(dict.middleEntry, coeffs)

	return &dict
}

/*func (dict *GaussianDictionary) GetAsMat(numEntries int) *mat.Matrix {
	dictionary := mat.NewDense(numEntries, numEntries, nil)
	for i := 0; i < numEntries; i++ {
		dictionary.Set(i, i, 1)
	}

	coeffVals := dict.gaussian.GetCoeffs(int(math.Ceil(dict.gaussian.sigma) * 5))
	coeffs := mat.NewVecDense(len(coeffVals), coeffVals)

	dictionary = ConvolveMatVec(dictionary, coeffs)

	fmt.Println(dictionary)

	return dictionary
}*/
