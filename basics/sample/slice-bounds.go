package sample

import "fmt"

func CallSliceBounds() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
