package sample

import (
	"fmt"
	"math"
)

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func main() {
// 	hypot := func(x, y float64) float64 {
// 		return math.Sqrt(x*x + y*y)
// 	}
// 	fmt.Println(hypot(5, 12))

// 	fmt.Println(compute(hypot))
// 	fmt.Println(compute(math.Pow))
// }

func (s *Sample) FunctionValus() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(s.compute(hypot))
	fmt.Println(s.compute(math.Pow))
}

func (s *Sample) compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
