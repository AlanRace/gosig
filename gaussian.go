package gosig

import "math"

type Gaussian struct {
	sigma float64

	scalar      float64
	denominator float64
}

func NewGaussian(sigma float64) *Gaussian {
	var gaussian Gaussian
	gaussian.SetSigma(sigma)

	return &gaussian
}

func (gaussian *Gaussian) Sigma() float64 {
	return gaussian.sigma
}

func (gaussian *Gaussian) SetSigma(sigma float64) {
	gaussian.sigma = sigma

	gaussian.scalar = 1 / (math.Sqrt(2*math.Pi) * gaussian.sigma)
	gaussian.denominator = 2 * gaussian.sigma * gaussian.sigma
}

func (gaussian *Gaussian) GetCoeffs(halfWidth int) []float64 {
	coeffs := make([]float64, halfWidth*2+1)

	for x := -halfWidth; x <= halfWidth; x++ {
		coeffs[x+halfWidth] = gaussian.scalar * math.Exp(-float64(x*x)/gaussian.denominator)
	}

	return coeffs
}

func FitGaussians(signal *Signal, gaussian *Gaussian, numGaussians, minimumDistance int) {

}
