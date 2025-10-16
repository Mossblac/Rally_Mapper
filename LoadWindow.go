package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadT() {

	/*loadWin.SetCloseIntercept(func() {
		loadWin.Hide()
	})*/
	Home := widget.NewButton("home", func() {
		HomeScreen()
	})
	CHome := container.NewCenter(Home)

	text := widget.NewLabel("Select Track to Load")
	text.Alignment = fyne.TextAlignCenter
	Ctext := container.NewCenter(text)

	cat, err := LoadCatalog("./Saves/catalog.json")
	if err != nil {
		fmt.Printf("error loading catalogue: %v\n", err)
	}
	catList := MakeCatalogList(cat, onSelect)

	bkg := canvas.NewImageFromFile("./images/menuBkG.jpg")

	centerStack := container.NewStack(bkg, catList)

	LoadListWithButton := container.NewBorder(Ctext, CHome, nil, nil, centerStack)

	mainWin.SetContent(LoadListWithButton)
}

func onSelect(T CatalogEntry) {
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
