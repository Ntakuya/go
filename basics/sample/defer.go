package sample

import "fmt"

func Defer() string {
	defer fmt.Println("world")

	fmt.Println("real hello")

	return "return just and called is latest because return value is function completed"
}
