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

var Obstacles []string
var TrackSize int
var ObDisplay string

func HomeScreen() {
	Loading = false
	Obstacles = nil

	var courseType string

	Track = Track[:0]
	RevCount = 0
	TrackSize = 0

	ObDisplayText := widget.NewLabel("")
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
	lowbridge.SetMinSize(fyne.NewSize(65, 65))
	parkingBlock := canvas.NewImageFromResource(ParkingBlock_Verticalicon)
	parkingBlock.FillMode = canvas.ImageFillContain
	parkingBlock.SetMinSize(fyne.NewSize(65, 65))
	curb_drop := canvas.NewImageFromResource(CurbDrop_Verticalicon)
	curb_drop.FillMode = canvas.ImageFillContain
	curb_drop.SetMinSize(fyne.NewSize(65, 65))

	iconBackground := canvas.NewRectangle(color.Black)
	iconBackground.SetMinSize(fyne.NewSize(70, 70))
	iconBackground.CornerRadius = 20

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
	selectbackground.CornerRadius = 20
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

	InfoB := widget.NewButton("Instructions", func() {
		showInstructions()
	})

	Warning := canvas.NewText("Select Track Size to Create", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	Warning.Alignment = fyne.TextAlignCenter
	Warning.TextStyle.Bold = true
	Warning.TextSize = 40
	Warning.Hide()

	createB := widget.NewButton("create", func() {
		Loading = false
		if TrackSize == 0 {
			Warning.Show()
			Warning.Refresh()
		} else {
			fyne.Do(func() {
				Grid_Widget(courseType, TrackSize, none)
			})
		}
	})

	loadTrack := widget.NewButton("Load Track", func() {
		LoadT()
	})

	CheckBox := container.NewHBox(SmallText, TtypeCheck)

	picktracksizeText := canvas.NewText("Select Track Size:", color.RGBA{R: 54, G: 1, B: 63, A: 255})
	picktracksizeText.TextStyle.Bold = true
	picktracksizeText.TextSize = 20

	TrackSizeSelectCentered := container.NewCenter(TrackSizeSelectStack)

	SelectObstacles := widget.NewLabel("Choose Obstacles\nMax 10")
	SelectObstacles.Alignment = fyne.TextAlignCenter
	SelObBKG := canvas.NewRectangle(color.Black)
	SelObBKG.CornerRadius = 20
	SelectObTxtStack := container.NewStack(SelObBKG, SelectObstacles)
	SelectObTextBox := container.NewHBox(layout.NewSpacer(), SelectObTxtStack, layout.NewSpacer())

	CheckAndSelectBox := container.NewVBox(CheckBox, picktracksizeText, TrackSizeSelectCentered, SelectObTextBox, IconVBox)

	buttonBox := container.NewHBox(InfoB, layout.NewSpacer(), loadTrack, layout.NewSpacer(), createB)

	centeredBox := container.NewCenter(buttonBox)

	buttonBoxWithWarning := container.NewVBox(Warning, centeredBox)

	ObstacleDisplayBKGRD := canvas.NewRectangle(color.Black)
	ObstacleDisplayBKGRD.CornerRadius = 20
	ObstacleDisplayBKGRD.SetMinSize(fyne.NewSize(170, 300))

	ObDisplayStack := container.NewStack(ObstacleDisplayBKGRD, container.NewCenter(ObDisplayText))
	obDisplayVbox := container.NewVBox(layout.NewSpacer(), ObDisplayStack, layout.NewSpacer())

	borderBox := container.NewBorder(container.NewCenter(LargeText), buttonBoxWithWarning, CheckAndSelectBox, obDisplayVbox)

	res := fyne.NewStaticResource("images/backgroundcropped.jpg", resourceBackgroundcroppedJpgData)
	bkg := canvas.NewImageFromResource(res)
	bkg.FillMode = canvas.ImageFillStretch

	res2 := fyne.NewStaticResource("images/rally_mapper_logo.png", resourceRallymapperlogoPngData)
	logo := canvas.NewImageFromResource(res2)
	logo.FillMode = canvas.ImageFillContain

	FullContent := container.NewStack(bkg, logo, borderBox)

	mainWin.SetContent(FullContent)

}

func showInstructions() {

	textLabel := widget.NewLabel(`
	Rally Mapper: 
	
	The default track style is a looping course,
	you can switch to a straight course with the toggle select.

	Select the size of your track - required - 
	This is the maximum area the course can occupy

	Select the obstacles that you want included on the track
	you can choose up to 10 that will be placed in the 
	order you selected. 

	Click "Create" to randomly generate a new Track!

	A Punch is a soft wireless pressure switch 
	they act as checkpoints throughout the track. 
	They are placed automatically
	and the time reduction for each is calculated for you. 

	Knock, hit, smash, tap, or punch 
	to obtain a small time reduction on your run. 

	The average track time, and specific time reductions for each Punch
	can be found by pressing the -Punch Info- button below the Track display

	Inspect the layout with the zoom buttons, and save your favorites

	The save button will open a small window that allows you to name your 
	Track and saves it under the name you provide. 

	Warning: each track has a unique ID, That is not based on the name you assign
	you can have multiple tracks under the same name or a track with no name at all. 

	To load, click the -load- button on the Home Menu 
	and select your track from the list. 



	`)

	homeButton := widget.NewButton("Home", func() {
		HomeScreen()
	})
	centeredLabel := container.NewCenter(textLabel)
	textScroll := container.NewVScroll(centeredLabel)

	centeredTextContent := container.NewBorder(nil, homeButton, nil, nil, textScroll)

	bkg := canvas.NewImageFromFile("./images/menuBkG.jpg")

	instructStack := container.NewStack(bkg, centeredTextContent)

	mainWin.SetContent(instructStack)

}

func ObstacleListBuilderADD(obtype string) {
	if len(Obstacles) < 10 {
		Obstacles = append(Obstacles, obtype)
		ObDisplay = strings.Join(Obstacles, "\n")
	}
	if len(strings.Split(ObDisplay, "\n")) == 10 {
		return
	}

}

func ObstacleListBuilderREMOVE(obtype string) {
	var NewObstacles []string
	obcount := 0
	for _, item := range Obstacles {
		if item == obtype {
			obcount++
		}
	}

	if len(Obstacles) == 1 && obcount == 1 {
		Obstacles = nil
		ObDisplay = ""
	}

	if len(Obstacles) > 1 && obcount == 1 {
		for _, item := range Obstacles {
			if item != obtype {
				NewObstacles = append(NewObstacles, item)
				Obstacles = NewObstacles
				ObDisplay = strings.Join(Obstacles, "\n")
			}
		}

	}
	if obcount == 0 {
		return

	} else {
		lastIndex := strings.LastIndex(ObDisplay, obtype+"\n")
		if lastIndex == -1 {
			return
		} else {
			ObDisplay = ObDisplay[:lastIndex] + ObDisplay[lastIndex+len(obtype+"\n"):]
			Obstacles = strings.Split(ObDisplay, "\n")
		}
	}

	fmt.Printf("%v\n", Obstacles) //still not fucking working
}
