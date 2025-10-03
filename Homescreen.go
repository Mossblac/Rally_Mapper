package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen() {

	var courseType string
	var Obstacles []string
	var selectedObstacle string
	Oblist := ""
	Track = Track[:0]
	RevCount = 0

	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	createB := widget.NewButton("create", func() {
		fmt.Print("Creating New Track\n\n")
		Grid_Widget(courseType, len(Obstacles))
	})

	TextWindow := widget.NewLabel("")
	MaxObWindow := widget.NewLabel("")

	TypeOptions := []string{"loop", "linear"}

	ObstaclesOption := []string{"Curb", "Drop", "Platform", "Low-Bridge", "Slalom"}

	setTypeOption := func(value string) {
		courseType = value
	}

	setObstacleOption := func(s string) {
		selectedObstacle = s
	}

	addButton := widget.NewButton("Add", func() {
		if len(Obstacles) < 10 {
			Obstacles = append(Obstacles, selectedObstacle)
			Oblist += selectedObstacle + "          " + "\n"
			TextWindow.SetText(Oblist)
		} else {
			MaxObWindow.SetText("max number of obstacles added!")
		}

	})

	courseTypeSelect := widget.NewSelect(TypeOptions, setTypeOption)
	ObstacleSelect := widget.NewSelect(ObstaclesOption, setObstacleOption)
	addObstacleBox := container.NewHBox(ObstacleSelect, addButton)
	selectBox := container.NewVBox(widget.NewLabel("Course Type"), courseTypeSelect, widget.NewLabel("# Obstacles"), addObstacleBox)
	centeredselectBox := container.NewCenter(selectBox)

	VBoxTextWindow := container.NewVBox(TextWindow, MaxObWindow)
	centeredTextWindow := container.NewCenter(VBoxTextWindow)

	buttonBox := container.NewHBox(mainB, createB)                                            //new horizontal box, making buttons side by side
	centeredBox := container.NewCenter(buttonBox)                                             // center the buttons and they will conform to standard size of text
	borderBox := container.NewBorder(nil, centeredBox, centeredselectBox, centeredTextWindow) // once they are centered, place them on the edge of the screen ( this order matters !)

	logo := canvas.NewImageFromFile("/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rally_mapper_logo.png")
	logo.FillMode = canvas.ImageFillContain

	centeredContent := container.NewStack(logo, borderBox)

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
