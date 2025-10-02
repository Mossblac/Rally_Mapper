package main

import (
	"time"

	"fyne.io/fyne/v2"
)

func DeterminePath_setStart(TrackType string, numRows, numCols int) {

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

	Track[0] = TrackCell{ //should assign first struct
		CurPosR: R,
		CurPosC: C,
		Visited: true,
		Start:   true,
	}

	/*for i := range Track {
		fmt.Printf("%+v\n", Track[i])
	}*/

	go func() {
		done1 := make(chan bool)
		done2 := make(chan bool)

		fyne.Do(func() {
			time.Sleep(1000 * time.Millisecond)
			SetStart(numRows, numCols)
			done1 <- true
		})
		<-done1

		fyne.Do(func() {
			time.Sleep(1000 * time.Millisecond)
			PickNext(numRows, numCols, 1)
			done2 <- true
		})
		<-done2

		fyne.Do(func() {
			time.Sleep(1000 * time.Millisecond)
			DisplayTrkImages()
		})
	}()
}
