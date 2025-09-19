package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// you are making the grid based on the inputs, so you need to create the widgets and collect the inputs first.

var mainWin fyne.Window
var course fyne.Window

func main() {

	Rally_Mapper := app.New()
	mainWin = Rally_Mapper.NewWindow("Rally Mapper")
	mainWin.Resize(fyne.NewSize(400, 300))
	course = Rally_Mapper.NewWindow("Course")

	HomeScreen()

	mainWin.ShowAndRun()

}
