package sample

import "fmt"

type pointerMethods struct {
}

func NewPointerMethod() *pointerMethods {
	return &pointerMethods{}
}

func (ps *pointerMethods) Pointer() {
	i, j := 42, 2701
	fmt.Println(i)  // value
	fmt.Println(&i) // pointer

	p := &i         // p is pointer of i
	fmt.Println(*p) // the value of pointer

	*p = 21 // insert value into p of pointer
	fmt.Println("*p is the value of p so, *p is", *p)
	fmt.Println("*p is the value of p so, p is", p)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
