package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Delete_Activator bool

func LoadT() {
	redLevelselect := &color.RGBA{R: 255, G: 0, B: 0, A: 25}
	redLevelConfirm := &color.RGBA{R: 255, G: 0, B: 0, A: 75}
	Del_Sel_rec := canvas.NewRectangle(redLevelselect)
	Del_Sel_rec.Hide()

	Del_Confirm_red := canvas.NewRectangle(redLevelConfirm)
	Del_Confirm_red.Hide()

	Home := widget.NewButton("home", func() {
		HomeScreen()
	})

	text := widget.NewLabel("Select Track to Load")
	text.Alignment = fyne.TextAlignCenter
	Ctext := container.NewCenter(text)

	DeleteBut := widget.NewButton("delete", func() {
		Delete_Activator = true
		text.Text = "Select Track to Delete"
		text.Refresh()
		Del_Sel_rec.Show()
	})

	Comfirm_text := widget.NewLabel("Confirm Delete Selected Track")
	text.Alignment = fyne.TextAlignCenter
	Cent_Confirm_text := container.NewCenter(Comfirm_text)

	Confirmed_DeleteBut := widget.NewButton("DELETE", func() {

	})
	CancelDeleteBut := widget.NewButton("Cancel", func() {

	})

	ConfirmButtons := container.NewHBox(CancelDeleteBut, Confirmed_DeleteBut)
	ConfirmBox := container.NewVBox(Cent_Confirm_text, ConfirmButtons)
	CentConfirmBox := container.NewCenter(ConfirmBox)
	CentConfirmBox.Hide()

	main_buttons := container.NewHBox(Home, DeleteBut)
	Cbuttons := container.NewCenter(main_buttons)

	cat, err := LoadCatalog("./Saves/catalog.json")
	if err != nil {
		fmt.Printf("error loading catalogue: %v\n", err)
	}
	catList := MakeCatalogList(cat, onSelect)

	resBK := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(resBK)

	centerStack := container.NewStack(bkg, Del_Sel_rec, Del_Confirm_red, catList, CentConfirmBox)

	LoadListWithButton := container.NewBorder(Ctext, Cbuttons, nil, nil, centerStack)

	mainWin.SetContent(LoadListWithButton)
}

func onSelect(T CatalogEntry) {
	if Delete_Activator {
		//need to find a way to have this show Del_Confirm_red and CentConfirmBox
	} else {
		Loading = true
		TrkPoint, err := LoadTrack("./Saves", T.ID)
		if err != nil {
			fmt.Printf("error loading Track: %v", err)
		}
		var trackT string
		if T.Ttype {
			trackT = "loop"
		} else {
			trackT = "linear"
		}
		Trk := *TrkPoint
		Grid_Widget(trackT, Trk.TSize, Trk)
	}
}

func DelConfirm(clBk *canvas.Rectangle, buttons *fyne.Container) {
	clBk.Show()
	buttons.Show()
}
