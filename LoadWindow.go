package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadT() {

	//del button
	//del mode window. hide
	//confirm window. hide to pass into DelT

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

	resBK := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(resBK)

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

func DelT() {
	// delete mode window (red square, message change).hide within LoadT, shows upon button press.
	//del confirm window, also hidden, pass into DelT to show

	//LoadT window with red partially transparent square on top of background
	//message changes to "select track to delete".

	//when selected, show the confirmation screen, "confirm delete Track : %v"
	//confirm screen has less transparent red background and small centered hbox with confirm and cancel buttons.

	//confirm will delete entry and return to LoadT, cancel will return to LoadT without deleting.
}
