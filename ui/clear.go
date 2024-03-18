package ui

import "runtime"

func ClearScreen() {
	os := runtime.GOOS
	if os == "linux" || os == "darwin" {
		print("\033[H\033[2J")
	} else if os == "windows" {
		print("\033[2J")
	}
}
