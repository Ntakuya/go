package sample

import (
	"runtime"
)

func SwitchSample() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X."
	case "linux":
		return "Linux"
	default:
		return os
	}
}
