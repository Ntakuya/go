package sample

import "fmt"

func ForIsGosWhile() int {
	sum := 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}
	return sum
}
