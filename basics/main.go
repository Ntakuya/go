package main

import (
	s "basics/sample"
	"fmt"
)

func main() {
	s.CallTypeAssertion()

	s.Do(21)
	s.Do("hello")
	s.Do(true)

	a := s.Person{"Arthur Dent", 42}
	z := s.Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

  if err := s.NewError(); err != nil {
    fmt.Println(err)
  }
}
