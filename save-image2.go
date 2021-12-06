package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// func main() {
// 	make(680, 493)
// }

func make(x int, y int) {
	robotgo.Sleep(3)

	bitmap := robotgo.CaptureScreen(x, y, 38, 50)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)

	fmt.Println("bitmap...", bitmap)
	img := robotgo.ToImage(bitmap)
	robotgo.SaveJpeg(img, "test_2.jpg")
}
