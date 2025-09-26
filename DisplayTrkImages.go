package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages(CellGrid [][]*fyne.Container) {
	for i := 1; i < len(Trk); i++ {
		fmt.Printf("%+v\n", Trk[i])
		pos, ok := Trk[i]["Position"].([]int)
		if ok {
			R := pos[0]
			C := pos[1]
			SetImageInCell(CellGrid, R, C, StartUP)

		}
	}
}
