package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DeterminePath_setStart(TrackType string, numRows, numCols int) {
	if len(Track) != numRows*numCols {
		if !VerifyTrackReset(Track) {
			fmt.Println("Track NOT RESET!")
		}
	}

	var R int
	var C int
	if TrackType == "loop" {
		TrackT = true
		R = numRows - 1
		C = 0
	}

	if TrackType == "linear" {
		TrackT = false
		R = numRows - 1
		C = 1
	}

	Track[0] = TrackCell{
		CurPosR: R,
		CurPosC: C,
		Visited: true,
		Start:   true,
	}
	resultCh := make(chan bool, 1)

	go func() {
		SetStart(numRows, numCols)
		ok := PickNext(numRows, numCols, 1)
		resultCh <- ok
	}()

	ok := <-resultCh
	fmt.Printf("ok = : %v\n", ok)

	if ok {
		fyne.Do(func() {
			DisplayTrkImages()
		})
	}
	//mainWin.Canvas().Capture() //this creates an image.Image of the completed map to use later.
}

func VerifyTrackReset(T []TrackCell) bool {
	for i := 0; i <= len(T); i++ {
		if T[i].CurPosR != -1 && T[i].CurPosC != -1 {
			return false
		}
	}
	return true
}
