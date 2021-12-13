package sample

import "fmt"

func (s *Sample) functionClosureAdd() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func (s *Sample) FunctionClosure() {
	pos, neg := s.functionClosureAdd(), s.functionClosureAdd()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
