package sample

import (
	"fmt"
	"math"
)

func PowWithElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return lim
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
