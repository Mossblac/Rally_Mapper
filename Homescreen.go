package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() {
	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	createB := widget.NewButton("create", func() {
		Grid_Widget()
	})

	buttonBox := container.New(layout.NewHBoxLayout(), mainB, createB)

	centeredContent := container.New(layout.NewCenterLayout(), buttonBox)

	mainWin.SetContent(centeredContent)

	//you will add new inputs here-
	// if inputs == nil, add error to errs []errors and return
}

//check all you input errors here

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
