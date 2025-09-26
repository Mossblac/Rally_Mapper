package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages(CellGrid [][]*fyne.Container) {
	for _, m := range Trk {
		for key, value := range m {
			fmt.Printf("Trk list: %v  %v\n\n", key, value)
		}
	}
}
