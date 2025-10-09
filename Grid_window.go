// go
package main

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var CellGrid [][]*fyne.Container

type squareGridLayout struct{ rows, cols int }

func (l *squareGridLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	if l.rows == 0 || l.cols == 0 || len(objs) == 0 {
		return
	}
	cell := float32(math.Min(
		float64(size.Width/float32(l.cols)),
		float64(size.Height/float32(l.rows)),
	))
	if cell < 1 {
		cell = 1
	}
	gridW := cell * float32(l.cols)
	gridH := cell * float32(l.rows)
	offX := (size.Width - gridW) / 2
	offY := (size.Height - gridH) / 2

	i := 0
	for r := 0; r < l.rows; r++ {
		for c := 0; c < l.cols; c++ {
			if i >= len(objs) {
				return
			}
			x := offX + float32(c)*cell
			y := offY + float32(r)*cell

			objs[i].Move(fyne.NewPos(x, y))
			objs[i].Resize(fyne.NewSize(cell, cell))

			if cont, ok := objs[i].(*fyne.Container); ok {
				if len(cont.Objects) >= 2 {
					border := cont.Objects[0]
					stack := cont.Objects[1]
					border.Move(fyne.NewPos(0, 0))
					border.Resize(fyne.NewSize(cell, cell))
					const m float32 = 1
					stack.Move(fyne.NewPos(m, m))
					stack.Resize(fyne.NewSize(cell-2*m, cell-2*m))
				}
			}
			i++
		}
	}
}

func (l *squareGridLayout) MinSize([]fyne.CanvasObject) fyne.Size {

	return fyne.NewSize(float32(l.cols), float32(l.rows))
}

func Grid_Widget(TrackType string, numObstacles int) {
	fmt.Print("Creating New Track\n\n")
	NumObstacles = numObstacles
	Track = nil
	TrackLength = 0
	PreSetCuls = nil
	var numCols, numRows int
	twiceObSize := float32(numObstacles * 2)

	if TrackType == "loop" {
		TrackT = true
		numCols = 3 * numObstacles
		numRows = 3 * numObstacles
	}
	if TrackType == "linear" {
		TrackT = false
		numCols = 3
		numRows = 3 * numObstacles
	}

	CellGrid = make([][]*fyne.Container, numRows)
	objs := make([]fyne.CanvasObject, 0, numRows*numCols)
	for r := 0; r < numRows; r++ {
		CellGrid[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {

			border := canvas.NewRectangle(color.Gray16{Y: 0x4000})
			content := canvas.NewRectangle(color.Black)
			content.SetMinSize(fyne.NewSize(40-twiceObSize, 40-twiceObSize))

			stack := container.NewStack()

			bg := canvas.NewRectangle(color.Black)
			stack.Add(bg)

			cellContainer := container.NewWithoutLayout(border, stack)

			CellGrid[r][c] = cellContainer
			objs = append(objs, cellContainer)
		}
	}

	// Grid container that auto-scales cells and centers them
	grid := container.New(&squareGridLayout{rows: numRows, cols: numCols}, objs...)

	// Centering wrappers no longer needed; grid self-centers via layout.
	// Keep your Home button and border layout.
	homeButton := widget.NewButton("Home", func() { HomeScreen() })
	RunAgainButton := widget.NewButton("Re-Generate", func() { Grid_Widget(TrackType, numObstacles) })

	buttonbox := container.NewVBox(homeButton, RunAgainButton)

	gridWithHomeB := container.NewBorder(nil, buttonbox, nil, nil, grid)
	mainWin.SetContent(gridWithHomeB)

	for i := 0; i < numRows*numCols; i++ {
		cell := TrackCell{
			TrackInt: -1,
			CurPosR:  -1,
			CurPosC:  -1,
			PrevMov:  "",
			Visited:  false,
			Start:    false,
		}
		Track = append(Track, cell)
	}
	DeterminePath_setStart(TrackType, numRows, numCols)
}
