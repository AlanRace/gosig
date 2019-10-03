package gosig

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Dictionary interface {
	GetAsMat(numEntries int) *mat.Dense
}

type GaussianDictionary struct {
	gaussian *Gaussian
}

func NewGaussianDictionary(gaussian *Gaussian) *GaussianDictionary {
	return &GaussianDictionary{gaussian: gaussian}
}

func (dict *GaussianDictionary) GetAsMat(numEntries int) *mat.Dense {
	dictionary := mat.NewDense(numEntries, numEntries, nil)
	for i := 0; i < numEntries; i++ {
		dictionary.Set(i, i, 1)
	}

	coeffVals := dict.gaussian.GetCoeffs(int(math.Ceil(dict.gaussian.sigma) * 5))
	coeffs := mat.NewVecDense(len(coeffVals), coeffVals)

	dictionary = ConvolveMatVec(dictionary, coeffs)

	return dictionary
}
