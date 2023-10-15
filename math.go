package GoUtil

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Min[N constraints.Integer | constraints.Float](x, y N) N {
	return N(math.Min(float64(x), float64(y)))
}
func Max[N constraints.Integer | constraints.Float](x, y N) N {
	return N(math.Max(float64(x), float64(y)))
}
