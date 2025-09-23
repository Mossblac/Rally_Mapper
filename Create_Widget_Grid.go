package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//startRightImage := canvas.NewImageFromFile("/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startRight.png")

func Grid_Widget(value string, numObstacles int) {
	var numCols int
	var numRows int
	twiceObSize := float32(numObstacles * 2)

	if value == "loop" {
		numCols = 3 * numObstacles
		numRows = 3 * numObstacles
	}

	if value == "linear" {
		numCols = 3
		numRows = 3 * numObstacles
	}

	grid := container.New(layout.NewGridLayout(numCols)) // cols is your desired number of columns

	// Store references to the cell containers
	cellContainers := make([][]*fyne.Container, numRows)
	for r := 0; r < numRows; r++ {
		cellContainers[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black)                              // Or any initial color
			rect.SetMinSize(fyne.NewSize((40 - twiceObSize), (40 - twiceObSize))) // Set a minimum size for the rectangle
			cellContainer := container.NewStack(rect)                             // Use MaxLayout to make the rect fill the container
			cellContainers[r][c] = cellContainer
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

	gridWithHomeB := container.New(layout.NewBorderLayout(nil, homeButton, nil, nil),
		homeButton,
		scrollingGrid,
	)

	mainWin.SetContent(gridWithHomeB)
}

func SetImageInCell(row, col int, image *canvas.Image) {

}
