package main

import (
	"basics/sample"
	"fmt"
)

func main() {
  fmt.Println(sample.ForIsGosWhile())
  fmt.Println(sample.Sqrt(2), sample.Sqrt(-4))
  fmt.Println(sample.Pow(3,2,10), sample.Pow(3, 3, 20))
  fmt.Println(sample.PowWithElse(3,2,10), sample.PowWithElse(3, 3, 20))
  fmt.Println(sample.SwitchSample())
  fmt.Println(sample.SwitchEvaluationOrder())
  fmt.Println(sample.SwitchWithNoCondition())
  fmt.Println(sample.Defer())
  sample.DeferMulti()
}
