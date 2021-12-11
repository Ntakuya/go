package sample

import "fmt"

func (s *Sample) CallMaps() {
	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex, 0)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
