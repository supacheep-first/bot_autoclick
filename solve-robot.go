package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

//		  	  x   y
// top right 788 443
// bot left 532 610
/*
x  <------->

y   ^
	|
	|
	|
	v
*/

func solveRobot() bool {
	robotgo.MilliSleep(2)
	_jigsawColor := []string{"959595", "636363", "939393", "7d7d7d", "7c7c7c", "646464", "9d9d9d", "adadad", "696969", "6c6c6c", "6d6d6d", "999999", "757575", "767676", "b0b0b0", "9c9c9c", "b1b1b1", "9f9f9f", "6a6a6a", "838383", "8d8d8d", "989898", "686868", "a6a6a6", "aaaaaa", "7b7b7b", "9e9e9e", "6e6e6e", "8b8b8b", "acacac", "a9a9a9", "929292", "9a9a9a", "878787", "979797", "5e5e5e", "707070", "4f4f4f", "666666", "515151", "8e8e8e", "7f7f7f", "606060", "5d5d5d", "8c8c8c", "535353", "858585", "949494", "a0a0a0", "808080", "afafaf", "868686", "919191", "a3a3a3", "b2b2b2", "747474", "a2a2a2", "8f8f8f", "5c5c5c", "595959", "909090", "787878", "777777", "626262", "848484", "aeaeae", "6f6f6f", "a8a8a8", "9b9b9b", "7e7e7e", "676767", "969696", "585858", "5f5f5f", "616161", "656565", "898989", "505050", "5b5b5b", "a5a5a5", "797979", "8a8a8a", "a7a7a7", "a4a4a4", "a1a1a1", "727272", "888888", "737373", "4e4e4e", "525252", "545454", "ababab", "7a7a7a", "717171", "575757", "818181", "6b6b6b", "828282", "5a5a5a", "555555", "565656"}
	var _positionJigsaw []int
	var _manyMatch float64 = 0
	var _manyMatchPos int = 0

	var sizeimage1 = 40
	var sizeimage2 = 45
	// robotgo.MouseToggle("down")
	for y := 443; y < 610; y += 5 {
		for x := 788; x > 532; x -= 5 {
			// robotgo.Move(x, y)
			color := robotgo.GetPixelColor(x, y)
			color1 := robotgo.GetPixelColor(x, y+1)
			color_1 := robotgo.GetPixelColor(x, y-1)
			if stringInSlice(color, _jigsawColor) && stringInSlice(color1, _jigsawColor) && stringInSlice(color_1, _jigsawColor) {
				fmt.Println("found jigsaw")
				_positionJigsaw = append(_positionJigsaw, x-sizeimage1)
				_positionJigsaw = append(_positionJigsaw, y)
				bitmap := robotgo.CaptureScreen(_positionJigsaw[0], _positionJigsaw[1], sizeimage2, 50)
				// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
				defer robotgo.FreeBitmap(bitmap)

				// fmt.Println("bitmap...", bitmap)
				img := robotgo.ToImage(bitmap)
				robotgo.SaveJpeg(img, "test_1.jpg")
				y = 999
				break
			}
		}
	}
	// robotgo.MouseToggle("up")
	//539 663

	fmt.Println(_positionJigsaw)

	_scroll := 0
	for x := 487; x < 800; x++ {
		color := robotgo.GetPixelColor(x, 663)
		if color == "ffdb23" {
			color10 := robotgo.GetPixelColor(x+10, 663)
			if color10 == "855540" {
				fmt.Println("found scroll bar : ", x, 663)
				_scroll = x
				break
			}
		}
	}

	// fmt.Println(_scroll)

	if _scroll != 0 {
		robotgo.Move(_scroll, 663)
		robotgo.Sleep(1)

		robotgo.MouseToggle("down")
		for x := _scroll; x < _positionJigsaw[0]+50; x++ {
			robotgo.MoveSmooth(x, 663)
			robotgo.MicroSleep(50)

			bitmap := robotgo.CaptureScreen(_positionJigsaw[0], _positionJigsaw[1], sizeimage2, 50)
			// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
			defer robotgo.FreeBitmap(bitmap)

			// fmt.Println("bitmap...", bitmap)
			img := robotgo.ToImage(bitmap)
			// var path = "test_" + strconv.Itoa(x) + ".jpg"
			// robotgo.SaveJpeg(img, path)
			robotgo.SaveJpeg(img, "test_2.jpg")

			_diffPercent := diffImage()
			if _diffPercent > _manyMatch {
				_manyMatch = _diffPercent
				_manyMatchPos = x
			}

		}

		_loopbtw1 := _manyMatchPos - 50
		_loopbtw2 := _manyMatchPos + 50

		var __manyMatch float64 = 0
		var __manyMatchPos int = 0
		for i := 0; i < 5; i++ {
			for j := _loopbtw1; j < _loopbtw2; j++ {
				robotgo.MoveSmooth(j, 663)
				robotgo.MicroSleep(50)

				bitmap := robotgo.CaptureScreen(_positionJigsaw[0], _positionJigsaw[1], sizeimage2, 50)
				// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
				defer robotgo.FreeBitmap(bitmap)

				// fmt.Println("bitmap...", bitmap)
				img := robotgo.ToImage(bitmap)
				// var path = "test_" + strconv.Itoa(x) + ".jpg"
				// robotgo.SaveJpeg(img, path)
				robotgo.SaveJpeg(img, "test_2.jpg")

				_diffPercent := diffImage()
				if _diffPercent > __manyMatch {
					__manyMatch = _diffPercent
					__manyMatchPos = j
					bitmap := robotgo.CaptureScreen(_positionJigsaw[0], _positionJigsaw[1], sizeimage2, 50)
					// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
					defer robotgo.FreeBitmap(bitmap)

					// fmt.Println("bitmap...", bitmap)
					img := robotgo.ToImage(bitmap)
					// var path = "test_" + strconv.Itoa(x) + ".jpg"
					// robotgo.SaveJpeg(img, path)
					robotgo.SaveJpeg(img, "super_match.jpg")
				}
			}
		}

		robotgo.Sleep(1)
		robotgo.MoveSmooth(__manyMatchPos-5, 663)
		robotgo.Sleep(1)

		robotgo.MouseToggle("up")
	}
	return true
}
