package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Grid_Widget() {

	numCols := 3
	numRows := 18

	gridItems := make([]fyne.CanvasObject, numCols*numRows)
	cells := make([][]*canvas.Rectangle, numRows)

	for r := 0; r < numRows; r++ {
		cells[r] = make([]*canvas.Rectangle, numCols)
		for c := 0; c < numCols; c++ {
			rect := canvas.NewRectangle(color.Black) // Initial white color
			rect.StrokeColor = color.White
			rect.StrokeWidth = 1
			cells[r][c] = rect
			gridItems[r*numCols+c] = rect

			// Add a tap gesture to change cell color on click

			rect.SetMinSize(fyne.NewSize(50, 50)) // Set a minimum size for visibility
			rect.Refresh()                        // Refresh to apply changes

			// Example of how to change a cell's color individually:
			// You would typically have a function or event handler to trigger this
			// For demonstration, let's change the color of a specific cell after a delay
			// go func(row, col int) {
			// 	time.Sleep(2 * time.Second)
			// 	cells[row][col].FillColor = color.RGBA{R: 255, A: 255} // Red
			// 	cells[row][col].Refresh()
			// }(r, c)
		}
	}

	grid := container.New(layout.NewGridLayoutWithColumns(numCols), gridItems...)

	// Example of changing a cell's color on interaction (e.g., button click)
	changeColorButton := widget.NewButton("Change Cell (1,1) to Blue", func() {
		if cells[1][1] != nil {
			cells[1][1].FillColor = color.RGBA{B: 255, A: 255} // Blue
			cells[1][1].Refresh()
		}
	})

	content := container.NewVBox(
		grid,
		changeColorButton,
	)
	course.SetContent(container.NewScroll(content))
	course.Resize(fyne.NewSize(200, 800))
	course.Show()
}
