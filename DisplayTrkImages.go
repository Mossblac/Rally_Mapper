package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages() {
	go func() {
		for i := 0; i < len(Track); i++ {
			m := Track[i]
			if m.CurPosR != -1 && m.CurPosC != -1 && !m.Start {

				fyne.Do(func() {

					icons := IconSet{
						Ic1: StraightUPicon,
						// Ic2, Ic3, Ic4 can be nil
					}

					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				})
			}
			time.Sleep(200 * time.Millisecond) // Wait between each image
		}

		fyne.Do(func() {
			fmt.Println("all images set")
		})
	}()
}
