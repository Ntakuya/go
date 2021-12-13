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
	v := s.Vertex{
		3, 4,
	}
	v.Scale(10)
	fmt.Println(vs.Abs())
	fmt.Println(v.Abs())
	s.Scale(&v, 10)
	fmt.Println(s.Abs(v))

	f := s.MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
