package main

import "github.com/go-vgo/robotgo"

func main() {
	// setup infinity loop
	// init time
	robotgo.Sleep(5)

	// idk why set but im don't forget set it!!!!
	robotgo.MouseSleep = 100

	infinity := 1
	for {
		// 	// re tab every 5 hr.
		// 	if infinity%5 == 0 {
		// 		doRefresh()
		// 	}
		// 	if infinity%5 == 0 || infinity == 1 {
		// 		if !login() {
		// 			doRefresh()
		// 			continue
		// 		}
		// 	}

		if !selectHero() {
			infinity = 0
			continue
		}

		bot()

		infinity++
	}

}

// func reOpenGame() {
// 	closeTab()
// 	openTabGame()
// }
