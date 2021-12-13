package main

import (
	s "basics/sample"
	"fmt"
	"math"
)

func main() {
	vs := &s.Vertex{
		3, 4,
	}
	fmt.Println(vs.Abs())
	fmt.Println(s.Abs(*vs))

	f := s.MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
