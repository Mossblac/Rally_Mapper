package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SaveWindow(numObstacles int) {
	SaveName := widget.NewEntry()
	SaveName.PlaceHolder = "Track Name"
	SaveName.MultiLine = false
	SaveName.Wrapping = fyne.TextWrapOff
	SaveName.OnChanged = func(name string) {
		TrackName = name
	}

	SaveButton := widget.NewButton("       Save        ", func() {
		save := TrackSave{
			Name:           TrackName,
			TimeToComplete: TotalTrackTime(),
			Ttype:          TrackT,
			TSize:          numObstacles,
			Punchlist:      PunchList,
			TrackData:      TrackClone,
		}
		Tsave := &save
		SaveNewTrack("./Saves", Tsave)
		ClickedOnce = true
		saveWin.Hide()
	})

	saveWin.SetCloseIntercept(func() {
		ClickedOnce = false
		saveWin.Hide()
	})

	SaveEntry := container.NewVBox(SaveName)
	SaveB := container.NewVBox(SaveButton)
	SaveEntryAndButton := container.NewVBox(SaveEntry, SaveB)

	resBK := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(resBK)

	SStack := container.NewStack(bkg, SaveEntryAndButton)

	Centered := container.NewBorder(nil, nil, nil, nil, SStack)

	saveWin.SetContent(Centered)
}

func SavedWindow() {
	close := widget.NewButton("close", func() {
		saveWin.Close()
	})

	saveWin.SetCloseIntercept(func() {
		saveWin.Hide()
	})

	AlreadySaved := widget.NewLabel("Track Saved")

	labelAndButton := container.NewHBox(AlreadySaved, close)

	saveWin.SetContent(labelAndButton)
}
