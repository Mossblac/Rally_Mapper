package main

import (
	"time"

	"fyne.io/fyne/v2"
)

var Trk []map[string]interface{}
var PossibleMoves []map[string][]int

func DeterminePath_setStart(TrackType string, CellGrid [][]*fyne.Container, numRows, numCols int) {

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

	//DisplayTrkImages(CellGrid)

	go func() {
		time.Sleep(1000 * time.Millisecond)
		SetStart(CellGrid, numRows, numCols)
		time.Sleep(1000 * time.Millisecond)
		PickNext(CellGrid, numRows, numCols, 1)
		time.Sleep(1000 * time.Millisecond)
		DisplayTrkImages(CellGrid)
	}()

}
