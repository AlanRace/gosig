package gosig

// Orthogonal matching pursuit

import (
	"math"

	"gonum.org/v1/gonum/blas/gonum"
	"gonum.org/v1/gonum/mat"
)

// OMP performs Orthognal Matching Pursuit to fit up to maxFits signals, separated by at least minDist, from the dictionary to the data
func OMP(signal *Signal, dict Dictionary, maxFits, minDist int) ([]int, []float64) {
	signalLength := signal.Len()
	dictionary := dict.GetAsMat(signalLength)

	data := mat.NewVecDense(signalLength, signal.Y)

	remainingDictionary := mat.DenseCopyOf(dictionary)

	residual := mat.VecDenseCopyOf(data)

	Z := mat.NewVecDense(signalLength, nil)
	yfit := mat.NewVecDense(signalLength, nil)

	squareNorm := mat.NewDense(1, 1, nil)
	squareNorm.Mul(data.T(), data)

	J := make([]int, signalLength)

	for i := 0; i < signalLength; i++ {
		J[i] = i
	}

	qual := make([]float64, maxFits)
	coeff := make([]float64, maxFits)

	optimalDictIndicies := make([]int, maxFits)

	var impl gonum.Implementation

	for k := 0; k < maxFits && remainingDictionary != nil; k++ {
		_, dictElements := remainingDictionary.Dims()

		if dictElements == 0 || len(J) == 0 {
			break
		}

		scaleProd := mat.NewVecDense(dictElements, nil)
		scaleProd.MulVec(remainingDictionary.T(), residual)

		maxVal := 0.0
		maxIndex := 0

		for i := 0; i < scaleProd.Len(); i++ {
			scaleProd.SetVec(i, math.Abs(scaleProd.AtVec(i)))

			if scaleProd.AtVec(i) > maxVal {
				maxVal = scaleProd.AtVec(i)
				maxIndex = i
			}
		}

		//fmt.Printf("Max Index %d (%f)\n", maxIndex, maxVal)

		optimalDictIndicies[k] = J[maxIndex]

		J = removeValues(J, J[maxIndex], minDist)

		P := mat.NewDense(signalLength, k+1, nil)
		for row := 0; row < signalLength; row++ {
			for col := 0; col <= k; col++ {
				P.Set(row, col, dictionary.At(row, optimalDictIndicies[col]))
			}
		}

		pp := mat.NewDense(k+1, k+1, nil)
		pp.Mul(P.T(), P)
		ppp := mat.NewDense(k+1, signalLength, nil)
		ppp.Solve(pp, P.T())

		tmp := mat.NewVecDense(k+1, nil)

		tmp.MulVec(ppp, data)

		coeff[k] = tmp.AtVec(k)

		Z.MulVec(P, tmp)
		Z.SubVec(Z, yfit)

		yfit.AddVec(yfit, Z)
		residual.SubVec(residual, Z)

		norm := impl.Dnrm2(yfit.Len(), yfit.RawVector().Data, 1)
		qual[k] = (norm * norm) / squareNorm.At(0, 0)

		remainingDictionary = stripDictionary(dictionary, J)
	}

	/*fmt.Println(optimalDictIndicies)
	fmt.Println(coeff)
	fmt.Println(qual)*/

	return optimalDictIndicies, qual
}

func stripDictionary(dictionary *mat.Dense, indiciesToKeep []int) *mat.Dense {
	if len(indiciesToKeep) == 0 {
		return nil
	}

	rows, _ := dictionary.Dims()

	newDictionary := mat.NewDense(rows, len(indiciesToKeep), nil)

	for index, indexToKeep := range indiciesToKeep {
		for row := 0; row < rows; row++ {
			newDictionary.Set(row, index, dictionary.At(row, indexToKeep))
		}
	}

	return newDictionary
}

func removeValues(slice []int, s int, dist int) []int {
	var result []int

	for _, value := range slice {
		if value < (s-dist) || value > (s+dist) {
			result = append(result, value)
		}
	}

	return result
}
