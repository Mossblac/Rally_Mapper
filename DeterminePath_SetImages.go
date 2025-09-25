package main

import (
	"time"

	"fyne.io/fyne/v2"
)

var Trk []map[string]interface{}
var PossibleMoves []map[string][]int

func DeterminePath_setimages(TrackType string, CellGrid [][]*fyne.Container, numRows, numCols int) {

	var R int
	var C int
	if TrackType == "loop" {
		R = numRows - 1
		C = 0
	}

	if TrackType == "linear" {
		R = numRows - 1
		C = 1
	}

	start := map[string]interface{}{
		"Position": []int{R, C},
		"TrkIndex": 0,
	}

	Trk = append(Trk, start)

	SetStart(CellGrid, numRows, numCols)
	go func() {
		time.Sleep(1 * time.Second)
		fyne.Do(func() {
			PickNext_Loop(numRows, numCols, 1)
		})
	}()

}
