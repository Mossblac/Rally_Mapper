package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Obstacles []string //use for ob list
var TrackSize int

func HomeScreen() {

	var courseType string

	Track = Track[:0]
	RevCount = 0
	TrackSize = 0

	courseType = "loop"
	Stroke := canvas.NewText("Loop", color.Black)
	Stroke.TextStyle.Bold = true
	Stroke.TextSize = 62

	LargeText := canvas.NewText("Loop", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	LargeText.TextStyle.Bold = true
	LargeText.TextSize = 60

	SmallText := canvas.NewText("Check to switch to : linear", color.Black)
	SmallText.TextSize = 20

	Linear_Checked_options := []string{"120ft", "180ft", "240ft", "300ft", "360ft", "420ft", "480ft"}
	Loop_UnChecked_options := []string{"3600sqft", "8100sqft", "14400sqft", "22500sqft", "32400sqft", "44100sqft", "57600sqft"}

	TrackSizeConverter := func(s string) {
		switch s {
		case "3600sqft":
			TrackSize = 2
		case "8100sqft":
			TrackSize = 3
		case "14400sqft":
			TrackSize = 4
		case "22500sqft":
			TrackSize = 5
		case "32400sqft":
			TrackSize = 6
		case "44100sqft":
			TrackSize = 7
		case "57600sqft":
			TrackSize = 8
		case "120ft":
			TrackSize = 2
		case "180ft":
			TrackSize = 3
		case "240ft":
			TrackSize = 4
		case "300ft":
			TrackSize = 5
		case "360ft":
			TrackSize = 6
		case "420ft":
			TrackSize = 7
		case "480ft":
			TrackSize = 8
		}
	}

	TrackSizeSelect := widget.NewSelect(Loop_UnChecked_options, TrackSizeConverter)

	TtypeCheck := widget.NewCheck("", func(checked bool) {
		if checked {
			TrackSizeSelect.Options = Linear_Checked_options
			TrackSizeSelect.ClearSelected()
			TrackSizeSelect.Refresh()
			courseType = "linear"
			LargeText.Text = "Linear"
			Stroke.Text = "Linear"
			SmallText.Text = "Check to switch to : loop"
			Stroke.Refresh()
			SmallText.Refresh()
			LargeText.Refresh()
		} else {
			TrackSizeSelect.Options = Loop_UnChecked_options
			TrackSizeSelect.ClearSelected()
			TrackSizeSelect.Refresh()
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
			Grid_Widget(courseType, TrackSize)
		})
	})

	TextWindow := widget.NewLabel("")
	MaxObWindow := widget.NewLabel("")

	CheckBoxV := container.NewHBox(TtypeCheck, SmallText)

	picktracksizeText := canvas.NewText("Select Track Size:", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	picktracksizeText.TextStyle.Bold = true
	picktracksizeText.TextSize = 20

	TrackSizeSelectCentered := container.NewCenter(TrackSizeSelect)
	CheckAndSelectBox := container.NewVBox(CheckBoxV, picktracksizeText, TrackSizeSelectCentered)

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
