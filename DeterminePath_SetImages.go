package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

var Trk []map[string]interface{}

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
		"visited":  true,
	}

	Trk = append(Trk, start)

	SetStart(numRows, numCols)

	//how to add image

	//SetImageInCell(CellGrid, R, C, StartRight)

	//start["image"] = "StartRightImage" // how to add image field

	fmt.Printf("%+v\n", Trk)

}
