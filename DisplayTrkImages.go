package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages(CellGrid [][]*fyne.Container) {
	for i := 1; i < len(Trk); i++ {
		m := Trk[i]
		pos, ok := m["Position"].([]int)
		if ok {
			SetImageInCell(CellGrid, pos[0], pos[1], RallyLogo)
		}
	}
	fmt.Println("all images set")
}
