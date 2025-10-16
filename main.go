package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var mainWin fyne.Window
var punchWin fyne.Window
var saveWin fyne.Window

func main() {

	Rally_Mapper := app.New()
	mainWin = Rally_Mapper.NewWindow("Rally Mapper")
	punchWin = Rally_Mapper.NewWindow("Punch Info")
	saveWin = Rally_Mapper.NewWindow("Save Track")
	saveWin.Resize(fyne.NewSize(100, 100))
	saveWin.Hide()
	punchWin.Resize(fyne.NewSize(300, 500))
	punchWin.Hide()
	mainWin.Resize(fyne.NewSize(800, 800))
	mainWin.SetMaster()

	HomeScreen()

	mainWin.ShowAndRun()

}

func init() {

}
