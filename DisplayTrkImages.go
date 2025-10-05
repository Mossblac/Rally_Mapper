package main

import (
	"fmt"
	"time"
)

func DisplayTrkImages() {

	go func() {

		for i := 0; i < len(Track); i++ {
			m := Track[i]
			if m.CurPosR != -1 && m.CurPosC != -1 {

				if !m.Rev {
					SetImageInCell(m.CurPosR, m.CurPosC, m.Image)
				}
				if m.Rev {
					icons := IconSet{}
					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				}
			}
			time.Sleep(50 * time.Millisecond) // Wait between each image
		}

		for i := range Track {
			fmt.Printf("%+v\n\n", Track[i])
		}
		fmt.Println("all images set")

	}()
}
