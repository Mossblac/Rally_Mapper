package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Obstacles []string //not resetting to nil when returning to homscreen, but you need to change the whole setup

func HomeScreen() {

	var courseType string
	var selectedObstacle string
	Oblist := ""
	Track = Track[:0]
	RevCount = 0

	courseType = "loop"
	Stroke := canvas.NewText("Loop", color.Black)
	Stroke.TextStyle.Bold = true
	Stroke.TextSize = 62

	LargeText := canvas.NewText("Loop", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	LargeText.TextStyle.Bold = true
	LargeText.TextSize = 60

	SmallText := canvas.NewText("Check to switch to : linear", color.Black)
	SmallText.TextSize = 20

	TtypeCheck := widget.NewCheck("", func(checked bool) {
		if checked {
			courseType = "linear"
			LargeText.Text = "Linear"
			Stroke.Text = "Linear"
			SmallText.Text = "Check to switch to : loop"
			Stroke.Refresh()
			SmallText.Refresh()
			LargeText.Refresh()
		} else {
			courseType = "loop"
			LargeText.Text = "Loop"
			Stroke.Text = "Loop"
			SmallText.Text = "Check to switch to : linear"
			Stroke.Refresh()
			SmallText.Refresh()
			LargeText.Refresh()
		}

	})
	StrokeStack := container.NewStack(Stroke, LargeText)
	CenteredCheckBoxAndText := container.NewCenter(StrokeStack)

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

	ObstaclesOption := []string{"Curb/Drop", "Parking Block", "Low-Bridge", "Slalom"}

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

	ObstacleSelect := widget.NewSelect(ObstaclesOption, setObstacleOption)
	addObstacleBox := container.NewHBox(ObstacleSelect, addButton)
	CheckBoxV := container.NewHBox(TtypeCheck, SmallText)
	selectBox := container.NewVBox(widget.NewLabel("# Obstacles"), addObstacleBox)
	centeredselectBox := container.NewCenter(selectBox)
	CheckAndSelectBox := container.NewVBox(CheckBoxV, centeredselectBox)

	VBoxTextWindow := container.NewVBox(TextWindow, MaxObWindow)
	centeredTextWindow := container.NewCenter(VBoxTextWindow)

	buttonBox := container.NewHBox(mainB, createB)                                                                //new horizontal box, making buttons side by side
	centeredBox := container.NewCenter(buttonBox)                                                                 // center the buttons and they will conform to standard size of text
	borderBox := container.NewBorder(CenteredCheckBoxAndText, centeredBox, CheckAndSelectBox, centeredTextWindow) // once they are centered, place them on the edge of the screen ( this order matters !)

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
