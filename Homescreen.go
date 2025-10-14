package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Obstacles []string //use for ob list
var TrackSize int
var ObDisplay string

func HomeScreen() {

	var courseType string

	Track = Track[:0]
	RevCount = 0
	TrackSize = 0

	ObDisplayText := widget.NewLabel("") //this is what goes in your display box as centered text
	ObDisplayText.Alignment = fyne.TextAlignCenter

	AddObstacleButtonLowBridge := widget.NewButton("+", func() {
		ObstacleListBuilderADD("LowBridge")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	RemoveObstacleButtonLowBride := widget.NewButton("-", func() {
		ObstacleListBuilderREMOVE("LowBridge")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	AddObstacleButtonParkingBlock := widget.NewButton("+", func() {
		ObstacleListBuilderADD("ParkingBlock")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	RemoveObstacleButtonParkingBlock := widget.NewButton("-", func() {
		ObstacleListBuilderREMOVE("ParkingBlock")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	AddObstacleButtonCurb_Drop := widget.NewButton("+", func() {
		ObstacleListBuilderADD("Curb_Drop")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	RemoveObstacleButtonCurb_Drop := widget.NewButton("-", func() {
		ObstacleListBuilderREMOVE("Curb_Drop")
		ObDisplayText.Text = ObDisplay
		ObDisplayText.Refresh()
	})

	lowbridge := canvas.NewImageFromResource(LowBridge_Verticalicon)
	lowbridge.FillMode = canvas.ImageFillContain
	lowbridge.SetMinSize(fyne.NewSize(40, 40))
	parkingBlock := canvas.NewImageFromResource(ParkingBlock_Verticalicon)
	parkingBlock.FillMode = canvas.ImageFillContain
	parkingBlock.SetMinSize(fyne.NewSize(40, 40))
	curb_drop := canvas.NewImageFromResource(CurbDrop_Verticalicon)
	curb_drop.FillMode = canvas.ImageFillContain
	curb_drop.SetMinSize(fyne.NewSize(65, 65))

	iconBackground := canvas.NewRectangle(color.Black)
	iconBackground.SetMinSize(fyne.NewSize(70, 70))

	lowBridgeStack := container.NewStack(iconBackground, container.NewCenter(lowbridge))
	parkingBlockStack := container.NewStack(iconBackground, container.NewCenter(parkingBlock))
	curb_dropStack := container.NewStack(iconBackground, container.NewCenter(curb_drop))

	lowbridgeWithButtons := container.NewHBox(layout.NewSpacer(), RemoveObstacleButtonLowBride, lowBridgeStack, AddObstacleButtonLowBridge, layout.NewSpacer())
	parkingblockWithButtons := container.NewHBox(layout.NewSpacer(), RemoveObstacleButtonParkingBlock, parkingBlockStack, AddObstacleButtonParkingBlock, layout.NewSpacer())
	curb_dropWithButtons := container.NewHBox(layout.NewSpacer(), RemoveObstacleButtonCurb_Drop, curb_dropStack, AddObstacleButtonCurb_Drop, layout.NewSpacer())

	IconVBox := container.NewVBox(lowbridgeWithButtons, parkingblockWithButtons, curb_dropWithButtons) // icons still not showing up.

	canvas.Refresh(lowbridge)
	canvas.Refresh(parkingBlock)
	canvas.Refresh(curb_drop)

	courseType = "loop"

	LargeText := canvas.NewText("Loop", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	LargeText.Alignment = fyne.TextAlignCenter
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
	selectbackground := canvas.NewRectangle(color.Black)
	TrackSizeSelectStack := container.NewStack(selectbackground, TrackSizeSelect)

	TtypeCheck := widget.NewCheck("", func(checked bool) {
		if checked {
			TrackSizeSelect.Options = Linear_Checked_options
			TrackSizeSelect.ClearSelected()
			TrackSizeSelect.Refresh()
			courseType = "linear"
			LargeText.Text = "Linear"
			SmallText.Text = "Check to switch to : loop"
			SmallText.Refresh()
			LargeText.Refresh()
		} else {
			TrackSizeSelect.Options = Loop_UnChecked_options
			TrackSizeSelect.ClearSelected()
			TrackSizeSelect.Refresh()
			courseType = "loop"
			LargeText.Text = "Loop"
			SmallText.Text = "Check to switch to : linear"
			SmallText.Refresh()
			LargeText.Refresh()
		}

	})

	mainB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	createB := widget.NewButton("create", func() {
		fyne.DoAndWait(func() {
			Grid_Widget(courseType, TrackSize)
		})
	})

	CheckBox := container.NewVBox(TtypeCheck, SmallText)

	picktracksizeText := canvas.NewText("Select Track Size:", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	picktracksizeText.TextStyle.Bold = true
	picktracksizeText.TextSize = 20

	TrackSizeSelectCentered := container.NewCenter(TrackSizeSelectStack)
	CheckAndSelectBox := container.NewVBox(CheckBox, picktracksizeText, TrackSizeSelectCentered, IconVBox) //left box

	buttonBox := container.NewHBox(mainB, createB)
	centeredBox := container.NewCenter(buttonBox)
	ObstacleDisplayBKGRD := canvas.NewRectangle(color.Black)
	ObstacleDisplayBKGRD.SetMinSize(fyne.NewSize(170, 300))

	ObDisplayStack := container.NewStack(ObstacleDisplayBKGRD, container.NewCenter(ObDisplayText))
	obDisplayVbox := container.NewVBox(layout.NewSpacer(), ObDisplayStack, layout.NewSpacer())

	borderBox := container.NewBorder(container.NewCenter(LargeText), centeredBox, CheckAndSelectBox, obDisplayVbox)

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

func ObstacleListBuilderADD(obtype string) {
	if len(Obstacles) < 10 {
		Obstacles = append(Obstacles, obtype)
		ObDisplay = strings.Join(Obstacles, "\n")
	} else {
		ObDisplay += "\nmax 10" //keeps adding this
	}
}

func ObstacleListBuilderREMOVE(obtype string) {
	lastIndex := strings.LastIndex(ObDisplay, obtype+"\n")
	if lastIndex == -1 {
		return
	} else {
		ObDisplay = ObDisplay[:lastIndex] + ObDisplay[lastIndex+len(obtype+"\n"):]
		Obstacles = strings.Split(ObDisplay, "\n")
	}

	fmt.Printf("%v\n", Obstacles) //not removing last item.
}
