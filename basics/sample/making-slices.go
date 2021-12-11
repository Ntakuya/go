package sample

import "fmt"

func MakingSlices() {
	a := make([]int, 5)
	makingSlicesPrintSlice("a", a)

	b := make([]int, 0, 5)
	makingSlicesPrintSlice("b", b)

	c := b[:2]
	makingSlicesPrintSlice("c", c)

	d := c[2:5]
	makingSlicesPrintSlice("d", d)
}

func makingSlicesPrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
