package main

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
)

var Trk []map[string]interface{}
var PossibleMoves []map[string][]int
var Wg sync.WaitGroup

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

	SetStart(CellGrid, numRows, numCols) //DeterminePath is in a go routine inside a fyne.DoAndWait()

	go func() {
		time.Sleep(500 * time.Millisecond)
		fyne.Do(func() { //Do allows for Picknext and anything within it, to change graphics
			PickNext(CellGrid, numRows, numCols, 1)
		})
	}()

}
