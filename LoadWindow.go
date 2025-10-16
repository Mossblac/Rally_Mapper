package main

import "fmt"

func LoadT() {

	loadWin.SetCloseIntercept(func() {
		loadWin.Hide()
	})
	cat, err := LoadCatalog("./Saves/catalog.json")
	if err != nil {
		fmt.Printf("error loading catalogue: %v\n", err)
	}
	catList := MakeCatalogList(cat, onSelect)

	loadWin.SetContent(catList)
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
	loadWin.Hide()

}
