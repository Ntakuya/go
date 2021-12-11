package sample

import "time"

func SwitchWithNoCondition() string {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    return "Good morning!"
  default:
    return "Good Evening!"
  }
}
