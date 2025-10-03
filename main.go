package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var mainWin fyne.Window

func main() {

	Rally_Mapper := app.New()
	mainWin = Rally_Mapper.NewWindow("Rally Mapper")
	mainWin.Resize(fyne.NewSize(800, 600))
	mainWin.SetMaster()

	HomeScreen()

	mainWin.ShowAndRun()

}

func init() {

}
