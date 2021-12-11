package sample

import "fmt"

func AppendSliceValue() {
	var s []int
	printAppendSliceValue(s)

	s = append(s, 0)
	printAppendSliceValue(s)

	s = append(s, 1)
	printAppendSliceValue(s)

	s = append(s, 2, 3, 4)
	printAppendSliceValue(s)
}

func printAppendSliceValue(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
