package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Obstacles []string

func HomeScreen() {

	var courseType string
	var selectedObstacle string
	Oblist := ""
	Track = Track[:0]
	RevCount = 0

	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	createB := widget.NewButton("create", func() {
		fyne.DoAndWait(func() {
			Grid_Widget(courseType, len(Obstacles))
		})
	})

	TextWindow := widget.NewLabel("")
	MaxObWindow := widget.NewLabel("")

	TypeOptions := []string{"loop", "linear"}

	ObstaclesOption := []string{"Curb/Drop", "Parking Block", "Low-Bridge", "Slalom"}

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
			MaxObWindow.SetText("max number\n of obstacles\n added!")
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

	bkg := canvas.NewImageFromFile("./images/backgroundcropped.jpg")
	bkg.FillMode = canvas.ImageFillStretch

	logo := canvas.NewImageFromFile("./images/rally_mapper_logo.png")
	logo.FillMode = canvas.ImageFillContain

	FullContent := container.NewStack(bkg, logo, borderBox)

	mainWin.SetContent(FullContent)

}

func showInstructions() {
	textLabel := widget.NewLabel("Enter number of punches\n\nEnter physical equipment avalable\n\nEnter number of desired obstacles\n\nEnter 0-4 to reprisent percentage\n of required obstacles\n (0 = 0%, 2 = 50%, 4 = 100%)\n\nEnter maximum size of\n area in square feet\n - default is unlimited\n\nverify info and confirm")
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

/*label := canvas.NewText("Hello, Fyne!", color.White)
label.TextSize = 24 */
