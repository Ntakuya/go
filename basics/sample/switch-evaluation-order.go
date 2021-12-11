package sample

import (
	"fmt"
	"time"
)

func SwitchEvaluationOrder() string {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		return "Today"
	case today + 1:
		return "Tomorrow"
	case today + 2:
		return "In Two days"
	default:
		return "Too far away"
	}
}
