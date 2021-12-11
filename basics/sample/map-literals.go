package sample

import "fmt"

func (s *Sample) MapLiterals() {
	type vertex struct {
		Lat, Long float64
	}

	m := make(map[string]vertex)

	m = map[string]vertex{
		"Bell Labs": vertex{
			40.68433, -74.39967,
		},
		"Google": vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}
