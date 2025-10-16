package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/canvas"
)

var TrackClone []TrackCell

func DisplayTrkImages(stop <-chan struct{}, track []TrackCell, working *canvas.Image) {
	working.Hide()
	go func() {
		select {
		case <-stop:
			return
		default:
		}

		TrackClone = CloneTrack(track)

		for i := 0; i < len(TrackClone); i++ {
			select {
			case <-stop:
				return
			default:
				m := TrackClone[i]
				if m.CurPosR != -1 && m.CurPosC != -1 {
					SetTrackImageInCell(m.CurPosR, m.CurPosC, m.Image)
				}
				select {
				case <-stop:
					return
				case <-time.After(50 * time.Millisecond):
				}
			}
		}

		for i := range TrackLength + 1 {
			fmt.Printf("%+v\n\n", Track[i])

		}
		fmt.Printf("Track Length : %v\n\n", TrackLength)
		fmt.Printf("Track Length with Corners: %v\n\n", len(Track))
		fmt.Printf("all images set\n\n")
		fmt.Printf("%v\n", TotalTrackTime())

	}()
}
