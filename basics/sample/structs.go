package sample

import "fmt"

func CallVertex() {
	v := Vertex{1, 2}
	fmt.Println(v)
}

func CallVertexX() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func CallWithPointer() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(p, *p)
	p.X = 1e9
	fmt.Println(v, p)
}

func CallWithDefference() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, p, *p)
}
