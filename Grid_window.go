// go
package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var CellGrid [][]*fyne.Container

// minimal custom layout: square cells, centered, fills available space
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
			i++
		}
	}
}

func (l *squareGridLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	// minimal footprint (one pixel per cell)
	return fyne.NewSize(float32(l.cols), float32(l.rows))
}

func Grid_Widget(TrackType string, numObstacles int) {
	var numCols, numRows int
	twiceObSize := float32(numObstacles * 2)

	if TrackType == "loop" {
		numCols = 3 * numObstacles
		numRows = 3 * numObstacles
	}
	if TrackType == "linear" {
		numCols = 3
		numRows = 3 * numObstacles
	}

	// Build cells once
	CellGrid = make([][]*fyne.Container, numRows)
	objs := make([]fyne.CanvasObject, 0, numRows*numCols)
	for r := 0; r < numRows; r++ {
		CellGrid[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize(40-twiceObSize, 40-twiceObSize))
			cellContainer := container.NewStack(rect)
			CellGrid[r][c] = cellContainer
			objs = append(objs, cellContainer)
		}
	}

	// Grid container that auto-scales cells and centers them
	grid := container.New(&squareGridLayout{rows: numRows, cols: numCols}, objs...)

	// Centering wrappers no longer needed; grid self-centers via layout.
	// Keep your Home button and border layout.
	homeButton := widget.NewButton("Home", func() { HomeScreen() })

	// If you want future zoom+scroll, wrap 'grid' in a scroll later.
	// For now, place grid directly.
	gridWithHomeB := container.NewBorder(nil, homeButton, nil, nil, grid)
	mainWin.SetContent(gridWithHomeB)

	// async path setup
	go func() {
		time.Sleep(1 * time.Second)
		fyne.Do(func() {
			DeterminePath_setStart(TrackType, numRows, numCols)
		})
	}()
}
