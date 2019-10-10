package gosig

// Signal stores pairs of values making up a signal.
// Signal implements the Sort Interface enabling the data to be sorted in ascending order based on X.
type Signal struct {
	X []float64
	Y []float64
}

func (pp Signal) Len() int {
	return len(pp.X)
}

func (pp Signal) Less(i, j int) bool {
	return pp.X[i] < pp.X[j]
}

func (pp Signal) Swap(i, j int) {
	tmpX := pp.X[i]
	pp.X[i] = pp.X[j]
	pp.X[j] = tmpX

	tmpY := pp.Y[i]
	pp.Y[i] = pp.Y[j]
	pp.Y[j] = tmpY
}

func Gradient(data []float64) []float64 {
	gradient := make([]float64, len(data))

	for i := 0; i < len(gradient); i++ {
		if i == 0 {
			gradient[i] = data[1] - data[0]
		} else if i == len(gradient)-1 {
			gradient[i] = data[len(gradient)-1] - data[len(gradient)-2]
		} else {
			gradient[i] = (data[i+1] - data[i-1]) / 2
		}
	}

	return gradient
}
