package series

import (
	"distclus/core"
	"distclus/real"
	"math"
	"strings"
)

// Space for processing reals of reals ([][]float64)
type Space struct {
	window     int
	innerSpace core.Space
}

// NewSpace create a new series space
func NewSpace(conf core.SpaceConf) Space {
	var innerSpace core.Space
	var sconf = conf.(Conf)
	switch strings.ToLower(sconf.InnerSpace) {
	default:
		innerSpace = real.NewSpace(conf)
	}
	return Space{
		window:     sconf.Window,
		innerSpace: innerSpace,
	}
}

func allocate(in [][]float64, dim int) (out [][]float64) {
	out = make([][]float64, dim)
	lenin := len(in[0])
	for index := range out {
		out[index] = make([]float64, lenin)
	}
	return
}

func interpolate(in [][]float64, dim int) (out [][]float64) {
	out = make([][]float64, dim, dim)
	lenin := len(in)
	var pos float64
	copy(out[0], in[0])
	copy(out[dim-1], in[lenin-1])
	for index := range out[1 : dim-1] {
		pos = float64(index * lenin / dim)
		var ceil = math.Ceil(pos)
		var p = pos - ceil
		iindex := int(ceil)
		i0 := in[iindex]
		i1 := in[iindex+1]
		for ii, ii0 := range i0 {
			ii1 := i1[ii]
			out[index][ii] = float64((ii0 + ii1) * p)
		}
	}
	return
}

func (space Space) getMatrix(elemt1, elemt2 core.Elemt) (matrix [][]float64) {
	var e1 = elemt1.([][]float64)
	var e2 = elemt2.([][]float64)

	cols := len(e1)
	rows := len(e2)

	matrix = make([][]float64, cols)

	for col := range matrix {
		matrix[col] = make([]float64, rows)
		for row := range matrix {
			matrix[col][row] = space.innerSpace.Dist(e1[col], e2[row])
		}
	}

	return
}

func (space Space) path(elemt1, elemt2 core.Elemt) (path []float64) {
	var e1 = elemt1.([][]float64)
	var e2 = elemt2.([][]float64)

	len1 := len(e1)
	len2 := len(e2)

	var cols = e1
	var rows = e2

	var w = space.window

	if len1 > (len2 + w) {
		cols = interpolate(e1, len2+w)
		rows = e2
	} else if len2 > (len1 + w) {
		cols = interpolate(e2, len1+w)
		rows = e1
	}

	// initialize the matrix
	var matrix = make([][]float64, len(cols)+1)

	var inf = math.Inf(0)

	for colIndex := range matrix {
		matrix[colIndex] = make([]float64, len(rows)+1)
		for rowIndex := range matrix[colIndex] {
			matrix[colIndex][rowIndex] = inf
		}
	}
	matrix[0][0] = 0

	var Dist = space.innerSpace.Dist

	var lenCols = len(matrix)
	var lenRows = len(matrix[0])
	w = int(math.Max(float64(w), math.Abs(float64(lenCols-lenRows))))

	// path = make([]float64, 0, lenRows+w)

	for colIndex := range matrix[1:] {
		for rowIndex := int(math.Max(1, float64(colIndex-w))); rowIndex < int(math.Min(float64(lenRows), float64(colIndex+w))); rowIndex++ {
			var insertion = matrix[colIndex-1][rowIndex]
			var deletion = matrix[colIndex][rowIndex-1]
			var match = matrix[colIndex-1][rowIndex-1]
			var dist = Dist(cols[0], rows[0])
			var cost = dist + math.Min(insertion, math.Min(deletion, match))
			matrix[colIndex][rowIndex] = cost
			path = append(path, cost)
		}
	}

	return
}

// Dist computes euclidean distance between two nodes
func (space Space) Dist(elemt1, elemt2 core.Elemt) (sum float64) {
	var path = space.path(elemt1, elemt2)

	for _, dist := range path {
		sum += dist
	}

	return
}

// Combine computes combination between two nodes
func (space Space) Combine(elemt1 core.Elemt, weight1 int, elemt2 core.Elemt, weight2 int) {
	var Combine = space.innerSpace.Combine

	Combine(elemt1, weight1, elemt2, weight2)
}

// Copy creates a copy of a vector
func (space Space) Copy(elemt core.Elemt) core.Elemt {
	var rv = elemt.([][]float64)
	var copied = make([][]float64, len(rv))
	for i := range copied {
		copied[i] = make([]float64, len(rv[i]))
		copy(copied[i], rv[i])
	}
	return copied
}

// Dim returns input data dimension
func (space Space) Dim(data []core.Elemt) (dim int) {
	if len(data) > 0 {
		series := data[0].([][]float64)
		if len(series) > 0 {
			dim = len(series[0])
		}
	}
	return
}
