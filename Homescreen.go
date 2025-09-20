package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() {
	var courseType string
	var numObstacles int

	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	createB := widget.NewButton("create", func() {
		Grid_Widget(courseType, numObstacles)
	})

	TypeOptions := []string{"loop", "linear"}
	ObstaclesOption := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	setTypeOption := func(value string) {
		courseType = value
	}

	setObstacleOption := func(value string) {
		ObstaclesInt, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("error converting to integer: %v", err)
		}

		numObstacles = ObstaclesInt
	}

	courseTypeSelect := widget.NewSelect(TypeOptions, setTypeOption)
	numObstacleSelect := widget.NewSelect(ObstaclesOption, setObstacleOption)
	selectBox := container.NewVBox(widget.NewLabel("Course Type"), courseTypeSelect, widget.NewLabel("# Obstacles"), numObstacleSelect)
	centeredselectBox := container.NewCenter(selectBox)

	buttonBox := container.NewHBox(mainB, createB)                             //new horizontal box, making buttons side by side
	centeredBox := container.NewCenter(buttonBox)                              // center the buttons and they will conform to standard size of text
	borderBox := container.NewBorder(nil, centeredBox, centeredselectBox, nil) // once they are centered, place them on the edge of the screen ( this order matters !)

	logo := canvas.NewImageFromFile("/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rally_mapper_logo.png")
	logo.FillMode = canvas.ImageFillContain

	centeredContent := container.NewStack(logo, borderBox)

	mainWin.SetContent(logo)

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
