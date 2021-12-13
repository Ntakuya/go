package main

import (
	s "basics/sample"
)

func main() {
	// vs := &s.Vertex{
	// 	3, 4,
	// }
	// v := s.Vertex{
	// 	3, 4,
	// }
	// v.Scale(10)
	// fmt.Println(vs.Abs())
	// fmt.Println(v.Abs())
	// s.Scale(&v, 10)
	// fmt.Println(s.Abs(v))

	// f := s.MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())

  s.CallTypeAssertion()

  s.Do(21)
	s.Do("hello")
	s.Do(true)

}
