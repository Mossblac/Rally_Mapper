package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

var RevInt int

func ReverseProtocol(numRows, numCols, I int) {

	RevCandidate := Track[I]
	if !RevCandidate.Cul && !RevCandidate.Rev {
		RevInt = I - 1
		fmt.Printf("RevInt: %v\n", RevInt)
		fmt.Printf("creating Cul @: %v, %v\n", RevCandidate.CurPosR, RevCandidate.CurPosC)
		prev := RevCandidate.PrevMov
		var NewPrev = ""
		fmt.Printf("Original Previous value: %v\n", prev)
		switch prev {
		case "UP":
			NewPrev = "Down"
		case "DUPLeft":
			NewPrev = "DDownRight"
		case "DUPRight":
			NewPrev = "DDownLeft"
		case "Right":
			NewPrev = "Left"
		case "Left":
			NewPrev = "Right"
		case "DDownRight":
			NewPrev = "DUPLeft"
		case "DDownLeft":
			NewPrev = "DUPRight"
		case "Down":
			NewPrev = "UP"
		}

		Track[I] = TrackCell{
			CurPosR: RevCandidate.CurPosR,
			CurPosC: RevCandidate.CurPosC,
			PrevMov: NewPrev,
			Visited: RevCandidate.Visited,
			Image:   RevCandidate.Image,
			Start:   RevCandidate.Start,
			Cul:     true,
			Rev:     false,
		}
		fmt.Printf("Previous swapped to: %v\n", NewPrev)

		go func() {
			fyne.Do(func() {
				PickNext(numRows, numCols, I)
			})
		}()
		return
	}

	if RevCandidate.Cul || RevCandidate.Rev {
		reverser := Track[RevInt]
		fmt.Printf("creating reverser @: %v, %v\n", reverser.CurPosR, reverser.CurPosC)
		prev := reverser.PrevMov
		var NPrev = ""
		fmt.Printf("Original Previous value: %v\n", prev)
		switch prev {
		case "UP":
			NPrev = "Down"
		case "DUPLeft":
			NPrev = "DDownRight"
		case "DUPRight":
			NPrev = "DDownLeft"
		case "Right":
			NPrev = "Left"
		case "Left":
			NPrev = "Right"
		case "DDownRight":
			NPrev = "DUPLeft"
		case "DDownLeft":
			NPrev = "DUPRight"
		case "Down":
			NPrev = "UP"
		}

		Track[I+1] = TrackCell{
			CurPosR: reverser.CurPosR,
			CurPosC: reverser.CurPosC,
			PrevMov: NPrev,
			Visited: reverser.Visited,
			Image:   reverser.Image,
			Start:   reverser.Start,
			Cul:     false,
			Rev:     true,
		}

		RevCount++

		go func() {
			fyne.Do(func() {
				RevInt--
				PickNext(numRows, numCols, I+1)
			})
		}()
		return
	}

}
