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
				SetImageInCell(m.CurPosR, m.CurPosC, m.Image)
			}
			time.Sleep(50 * time.Millisecond) // Wait between each image
		}

		for i := range TrackLength + 1 {
			fmt.Printf("%+v\n\n", Track[i])
		}
		fmt.Printf("Track Length : %v\n", TrackLength)
		fmt.Printf("all images set\n\n")

	}()
}
