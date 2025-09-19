package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// you are making the grid based on the inputs, so you need to create the widgets and collect the inputs first.

var mainWin fyne.Window

func main() {
	Rally_Mapper := app.New()
	mainWin = Rally_Mapper.NewWindow("Rally Mapper")

	HomeScreen()

	mainWin.Resize(fyne.NewSize(400, 300))
	mainWin.ShowAndRun()
}

func HomeScreen() {
	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	centeredContent := container.New(layout.NewCenterLayout(), mainB)

	mainWin.SetContent(centeredContent)
}

func showInstructions() {
	textLabel := widget.NewLabel("Enter number of punches\nEnter physical equipment avalable\nEnter number of desired obstacles\nEnter 0-4 to reprisent percentage of required obstacles (0 = 0%, 2 = 50%, 4 = 100%)\nEnter maximum size of area in square feet - default is unlimited\nverify info and confirm")
	textLabel.Alignment = fyne.TextAlignCenter

	homeButton := widget.NewButton("Home", func() {
		HomeScreen()
	})

	textLayout := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		textLabel,
		layout.NewSpacer(),
	)

	centeredTextContent := container.New(layout.NewBorderLayout(nil, homeButton, nil, nil),
		homeButton,
		textLayout,
	)

	mainWin.SetContent(centeredTextContent)
}
