package main

import (
	"image/color"
	"time"

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

	var CellGrid = make([][]*fyne.Container, numRows)
	for r := 0; r < numRows; r++ {
		CellGrid[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize((40 - twiceObSize), (40 - twiceObSize))) // zoom in and out buttons will change this and refresh
			cellContainer := container.NewStack(rect)
			CellGrid[r][c] = cellContainer
			grid.Add(cellContainer)
		}
	}

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

	// you have to put the sleep within a go function for it to separate properly

	time.Sleep(1 * time.Second)
	fyne.DoAndWait(func() {
		DeterminePath_setStart(TrackType, CellGrid, numRows, numCols)
	})

}
