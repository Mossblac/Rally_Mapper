package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func PunchInfo() {

	CloseButton := widget.NewButton("Close", func() {
		punchWin.Hide()
	})

	CloseButtonCentered := container.NewHBox(layout.NewSpacer(), CloseButton, layout.NewSpacer())

	iconBackground := canvas.NewRectangle(color.Black)
	iconBackground.SetMinSize(fyne.NewSize(70, 70))

	Punch1 := canvas.NewImageFromResource(Punch1icon)
	Punch1.FillMode = canvas.ImageFillContain
	Punch1.SetMinSize(fyne.NewSize(65, 65))

	Punch2 := canvas.NewImageFromResource(Punch2icon)
	Punch2.FillMode = canvas.ImageFillContain
	Punch2.SetMinSize(fyne.NewSize(65, 65))

	Punch3 := canvas.NewImageFromResource(Punch3icon)
	Punch3.FillMode = canvas.ImageFillContain
	Punch3.SetMinSize(fyne.NewSize(65, 65))

	Punch4 := canvas.NewImageFromResource(Punch4icon)
	Punch4.FillMode = canvas.ImageFillContain
	Punch4.SetMinSize(fyne.NewSize(65, 65))

	Punch5 := canvas.NewImageFromResource(Punch5icon)
	Punch5.FillMode = canvas.ImageFillContain
	Punch5.SetMinSize(fyne.NewSize(65, 65))

	Punch6 := canvas.NewImageFromResource(Punch6icon)
	Punch6.FillMode = canvas.ImageFillContain
	Punch6.SetMinSize(fyne.NewSize(65, 65))

	Punch7 := canvas.NewImageFromResource(Punch7icon)
	Punch7.FillMode = canvas.ImageFillContain
	Punch7.SetMinSize(fyne.NewSize(65, 65))

	Punch8 := canvas.NewImageFromResource(Punch8icon)
	Punch8.FillMode = canvas.ImageFillContain
	Punch8.SetMinSize(fyne.NewSize(65, 65))

	Punch9 := canvas.NewImageFromResource(Punch9icon)
	Punch9.FillMode = canvas.ImageFillContain
	Punch9.SetMinSize(fyne.NewSize(65, 65))

	Punch10 := canvas.NewImageFromResource(Punch10icon)
	Punch10.FillMode = canvas.ImageFillContain
	Punch10.SetMinSize(fyne.NewSize(65, 65))

	Punch1stack := container.NewStack(iconBackground, container.NewCenter(Punch1))
	/*Punch2stack := container.NewStack(iconBackground, container.NewCenter(Punch2))
	Punch3stack := container.NewStack(iconBackground, container.NewCenter(Punch3))
	Punch4stack := container.NewStack(iconBackground, container.NewCenter(Punch4))
	Punch5stack := container.NewStack(iconBackground, container.NewCenter(Punch5))
	Punch6stack := container.NewStack(iconBackground, container.NewCenter(Punch6))
	Punch7stack := container.NewStack(iconBackground, container.NewCenter(Punch7))
	Punch8stack := container.NewStack(iconBackground, container.NewCenter(Punch8))
	Punch9stack := container.NewStack(iconBackground, container.NewCenter(Punch9))
	Punch10stack := container.NewStack(iconBackground, container.NewCenter(Punch10))*/

	P1WithText := container.NewHBox(layout.NewSpacer(), Punch1stack, labelFromPtime(PTimes(0)), layout.NewSpacer())

	AllInfoVBox := container.NewVBox(LabelFromTotalTime(TotalTrackTime()), P1WithText, layout.NewSpacer())

	PBorderBox := container.NewBorder(nil, CloseButtonCentered, AllInfoVBox, nil, nil)

	punchWin.SetContent(PBorderBox)
}

func PTimes(I int) (ptime time.Duration) {
	if PunchList == nil {
		return time.Duration(0.0)
	}
	if len(PunchList) >= I {
		return Track[PunchList[I]].PTime
	}
	return time.Duration(0.0)
}

func labelFromPtime(ptime time.Duration) *widget.Label {
	ptimeText := ptime.String()
	text := widget.NewLabel(ptimeText)
	text.Alignment = fyne.TextAlignCenter
	return text
}

func LabelFromTotalTime(Ttime time.Duration) *widget.Label {
	TtimeText := Ttime.String()
	text := widget.NewLabel("Average Track Time:  " + TtimeText)
	text.Alignment = fyne.TextAlignCenter
	return text
}
