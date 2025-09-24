package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Grid_Widget(TrackType string, numObstacles int) {
	var numCols int
	var numRows int
	twiceObSize := float32(numObstacles * 2)

	if TrackType == "loop" {
		numCols = 3 * numObstacles
		numRows = 3 * numObstacles
	}

	if TrackType == "linear" {
		numCols = 3
		numRows = 3 * numObstacles
	}

	grid := container.New(layout.NewGridLayout(numCols))

	cellContainers := make([][]*fyne.Container, numRows)
	for r := 0; r < numRows; r++ {
		cellContainers[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize((40 - twiceObSize), (40 - twiceObSize))) // zoom in and out buttons will change this and refresh
			cellContainer := container.NewStack(rect)
			cellContainers[r][c] = cellContainer
			grid.Add(cellContainer)
		}
	}

	DeterminePath_setimages(TrackType, cellContainers, numRows, numCols)

	cgrid := container.NewCenter(grid)
	centeredGrid := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		cgrid,
		layout.NewSpacer(),
	)
	scrollingGrid := container.NewScroll(centeredGrid)

	homeButton := widget.NewButton("Home", func() {
		HomeScreen()
	})

	gridWithHomeB := container.New(layout.NewBorderLayout(nil, homeButton, nil, nil), // create zoom in and out buttons, place in NewVBox, add to right side here
		homeButton,
		scrollingGrid,
	)

	mainWin.SetContent(gridWithHomeB)
}

func SetImageInCell(CellGrid [][]*fyne.Container, row, col int, imageName_Path []string) {
	if path, ok := ImagePaths[imageName_Path[0]]; ok {
		path.FillMode = canvas.ImageFillContain
		CellGrid[row][col].Objects = []fyne.CanvasObject{path}
		CellGrid[row][col].Refresh()
	} else {
		image := canvas.NewImageFromFile(imageName_Path[1])
		ImagePaths[imageName_Path[0]] = image
		image.FillMode = canvas.ImageFillContain
		CellGrid[row][col].Objects = []fyne.CanvasObject{image}
		CellGrid[row][col].Refresh()
	}
}

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
	var Trk []map[string]interface{}

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
