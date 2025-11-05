package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Delete_Activator bool
var ToDelete CatalogEntry

func LoadT() {
	Delete_Activator = false
	redLevelselect := &color.RGBA{R: 255, G: 0, B: 0, A: 25}
	Del_Sel_rec := canvas.NewRectangle(redLevelselect)
	Del_Sel_rec.Hide()

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

	main_buttons := container.NewHBox(Home, DeleteBut)
	Cbuttons := container.NewCenter(main_buttons)

	cat, err := LoadCatalog("./Saves/catalog.json")
	if err != nil {
		fmt.Printf("error loading catalogue: %v\n", err)
	}
	catList := MakeCatalogList(cat, onSelect)

	resBK := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(resBK)

	centerStack := container.NewStack(bkg, Del_Sel_rec, catList)

	LoadListWithButton := container.NewBorder(Ctext, Cbuttons, nil, nil, centerStack)

	mainWin.SetContent(LoadListWithButton)
}

func onSelect(T CatalogEntry) {
	if Delete_Activator {
		ConfirmWin(T)
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

func ConfirmWin(ToDelete CatalogEntry) {
	redLevelConfirm := &color.RGBA{R: 255, G: 0, B: 0, A: 75}
	Del_Confirm_red := canvas.NewRectangle(redLevelConfirm)

	TText := fmt.Sprintf("Confirm Delete Track: %v ", ToDelete.Name)

	Comfirm_text := widget.NewLabel(TText)
	Comfirm_text.Alignment = fyne.TextAlignCenter
	textBkg := canvas.NewRectangle(color.Black)
	textBkg.CornerRadius = 20
	textStack := container.NewStack(textBkg, Comfirm_text)
	Cent_Confirm_text := container.NewCenter(textStack)

	Confirmed_DeleteBut := widget.NewButton("DELETE", func() {
		cat, err := LoadCatalog("./Saves/catalog.json")
		if err != nil {
			fmt.Printf("error loading catalogue: %v\n", err)
		}
		var NewCat Catalog
		for _, entry := range cat.Entries {
			if entry.ID != ToDelete.ID {
				NewCat.Add(entry)
			}
		}
		catalogPath := filepath.Join("./Saves", "catalog.json")
		atomicWriteJSON(catalogPath, &NewCat, 0o644)

		trackPath := filepath.Join("./Saves", "tracks", ToDelete.ID+".json")
		err = os.Remove(trackPath)
		if err != nil {
			fmt.Printf("error deleting track JSON: %v", err)
		}

		fmt.Printf("Track: %v deleted", ToDelete.Name)
		LoadT()
	})
	CancelDeleteBut := widget.NewButton("Cancel", func() {
		LoadT()
	})

	ConfirmButtons := container.NewHBox(CancelDeleteBut, Confirmed_DeleteBut)
	CconfirmButtons := container.NewCenter(ConfirmButtons)
	ConfirmBox := container.NewVBox(Cent_Confirm_text, CconfirmButtons)
	CentConfirmBox := container.NewCenter(ConfirmBox)

	resBK := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(resBK)

	centerStack := container.NewStack(bkg, Del_Confirm_red, CentConfirmBox)

	mainWin.SetContent(centerStack)

}
