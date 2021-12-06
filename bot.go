package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func bot() {

	robotgo.MouseSleep = 100
	robotgo.Sleep(5)
	robotgo.Move(637, 530) //treasure mode
	robotgo.Click("left")
	fmt.Println("click treasure mode")
	robotgo.Sleep(10)

	for i := 0; i < 60; i++ {
		robotgo.Move(653, 743) // next map
		robotgo.Sleep(2)
		robotgo.Click("left")
		robotgo.Sleep(58) // 1min
		// robotgo.Sleep(6) // 6s

		checkRobot()
	}

	// last click next map
	robotgo.Click("left")
	robotgo.Sleep(10)

	robotgo.Move(217, 277) //back to home
	robotgo.Click("left")
	fmt.Println("click back to home")
	robotgo.Sleep(10)
}
