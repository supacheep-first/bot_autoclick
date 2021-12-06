package main

import (
	"github.com/go-vgo/robotgo"
)

func checkRobot() bool {
	for i := 0; i < 10; i++ {
		topcolors := getlineColors(530, 783, 436)
		botcolors := getlineColors(530, 783, 613)
		if compareColor(topcolors, "b57052") && compareColor(botcolors, "b57052") {
			if !solveRobot() {
				return false
			}
			robotgo.Sleep(10)
		} else {
			break
		}
	}
	return true
}
