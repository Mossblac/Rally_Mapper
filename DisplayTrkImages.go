package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages() {
	go func() {
		for i := 1; i < len(Trk); i++ {
			m := Trk[i]
			pos, ok := m["Position"].([]int)
			if ok {
				fyne.Do(func() {
					// Remove the delay from SetImageInCell, handle timing here
					SetImageInCell(pos[0], pos[1], RallyLogo, 0)
				})
				time.Sleep(200 * time.Millisecond) // Wait between each image
			}
		}
		fyne.Do(func() {
			fmt.Println("all images set")
		})
	}()
}
