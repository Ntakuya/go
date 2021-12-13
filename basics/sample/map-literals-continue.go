package sample

import "fmt"

func (s *Sample) MapLiteralsContinue() {
	type vertex struct {
		Lat, Long float64
	}

	m := map[string]vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}
