package sample

import "fmt"

func PrinNilSlices() {
	var s []int
	m := make([]int, 0)
	// var a [0]int
	printSliceWithCapAndLength(s)
	printSliceWithCapAndLength(m)
	// printSliceWithCapAndLength(a)
}

func printSliceWithCapAndLength(s []int) {
	fmt.Println(s, cap(s), len(s))
	if s == nil {
		fmt.Println("NILL!!!")
	}
}
