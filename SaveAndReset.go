package main

import "fyne.io/fyne/v2"

func ResetAndTryAgain(numRows, numCols int) {
	Track = nil

	for i := 0; i < numRows*numCols; i++ {
		cell := TrackCell{
			CurPosR: -1,
			CurPosC: -1,
			PrevMov: "",
			Visited: false,
			Start:   false,
		}
		Track = append(Track, cell)
	}

	if TrackT {
		go func() {
			fyne.Do(func() {
				DeterminePath_setStart("loop", numRows, numCols)
			})
		}()
		return
	} else {
		go func() {
			fyne.Do(func() {
				DeterminePath_setStart("linear", numRows, numCols)
			})
		}()
		return
	}
}
