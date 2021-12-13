package main

import (
	s "basics/sample"
	"fmt"
)

func main() {
	vs := &s.Vertex{
    3, 4,
  }
	fmt.Println(vs.Abs())
}
