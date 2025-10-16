package main

import (
	"fmt"
	"time"
)

func DisplayTrkImages(stop <-chan struct{}) {

	go func() {
		select {
		case <-stop:
			return
		default:
		}

		trackClone := CloneTrack(Track)

		for i := 0; i < len(trackClone); i++ {
			select {
			case <-stop:
				return
			default:
				m := trackClone[i]
				if m.CurPosR != -1 && m.CurPosC != -1 {
					SetImageInCell(m.CurPosR, m.CurPosC, m.Image)
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
