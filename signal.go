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
