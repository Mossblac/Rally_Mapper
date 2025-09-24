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

	//how to add image

	SetImageInCell(CellGrid, R, C, StartRight)
	SetImageInCell(CellGrid, R, C+1, StartRight)
	SetImageInCell(CellGrid, R-1, C+1, StartAngleUR)

	start["image"] = "StartRightImage" // how to add image field

	fmt.Printf("%+v\n", Trk)

	/*if greater than 0, and less than numRows/Cols, and visited is false, add to list
	then randomly pick from that list.

	whichever is picked, create a new map named the picked coordinates AND set The "Position" value to the same.
	in that new map set "previous" to the previous cell coordinates( the first will just be: R, C)

	if previous == R+1, C-1 then previous image is set to face TopRightCorner and so on....
	*/

}
