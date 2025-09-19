package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// you are making the grid based on the inputs, so you need to create the widgets and collect the inputs first.

var mainWin fyne.Window
var course fyne.Window

func main() {
	test_cell := Cell{
		has_left_wall:   true,
		has_right_wall:  true,
		has_top_wall:    true,
		has_bottom_wall: true,
		x1:              1,
		y1:              1,
		cell_size:       5,
		window:          course,
		visited:         false,
	}

	Rally_Mapper := app.New()
	mainWin = Rally_Mapper.NewWindow("Rally Mapper")
	mainWin.Resize(fyne.NewSize(400, 300))
	course = Rally_Mapper.NewWindow("Course")

	HomeScreen()
	CT, CL, CB, CR, err := test_cell.Draw()
	if err != nil {
		fmt.Println("error creating lines")
	}

	courseContainer := container.NewWithoutLayout(CT, CL, CB, CR)
	course.SetContent(courseContainer)
	course.Resize(fyne.NewSize(float32(test_cell.cell_size*10), float32(test_cell.cell_size*10)))

	course.Show()

	mainWin.ShowAndRun()

}
