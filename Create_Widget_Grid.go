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

	grid := container.New(layout.NewGridLayout(numCols))

	cellContainers := make([][]*fyne.Container, numRows)
	for r := 0; r < numRows; r++ {
		cellContainers[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize((40 - twiceObSize), (40 - twiceObSize)))
			cellContainer := container.NewStack(rect)
			cellContainers[r][c] = cellContainer
			grid.Add(cellContainer)
		}
	}

	DeterminePath(cellContainers, numRows, numCols)

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

func SetImageInCell(CellGrid [][]*fyne.Container, row, col int, image *canvas.Image) {
	image.FillMode = canvas.ImageFillContain
	CellGrid[row][col].Objects = []fyne.CanvasObject{image}
	CellGrid[row][col].Refresh()
}

func DeterminePath(CellGrid [][]*fyne.Container, numRows, numCols int) {
	R := numRows - 1
	C := 0
	var Trk []map[string]interface{}

	start := map[string]interface{}{
		"Position": []int{R, C},
	}

	Trk = append(Trk, start)

	//how to add image
	StartRightImage := canvas.NewImageFromFile("/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startRight.png")
	SetImageInCell(CellGrid, R, C, StartRightImage)

	fmt.Printf("%+v\n", Trk)

	//how to search your structs

}
