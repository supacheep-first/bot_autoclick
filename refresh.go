package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func doRefresh() {
	for i := 0; i < 3; i++ {
		robotgo.Sleep(3)
		robotgo.KeyDown("lctrl")
		robotgo.KeyTap("f5")
		robotgo.KeyUp("lctrl")
	}
	fmt.Println("refresh page.")
}
