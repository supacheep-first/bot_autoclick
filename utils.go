package main

import "github.com/go-vgo/robotgo"

func getlineColors(xl, xr, y int) []string {
	var res []string

	for xl < xr {
		res = append(res, robotgo.GetPixelColor(xl, y))
		xl += 10
	}

	return res
}

func compareColors(colors []string, color ...string) bool {
	for _, _color := range colors {
		if contains(color, _color) {
			return true
		}
	}
	return false
}

func compareColor(colors []string, color string) bool {
	for _, _color := range colors {
		if color != _color {
			return false
		}
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
