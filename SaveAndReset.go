package main

import (
	"fyne.io/fyne/v2"
)

func ResetAndTryAgain() {
	RevCount = 0
	if TrackT {
		fyne.Do(func() {
			Grid_Widget("loop", NumObstacles)
		})
	} else {
		fyne.Do(func() {
			Grid_Widget("linear", NumObstacles)
		})
	}
}
