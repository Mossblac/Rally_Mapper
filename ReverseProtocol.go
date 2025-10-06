package main

import (
	"fmt"
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

		FirstCulImage := IconSet{
			Ic1: DetermineFirstCulIcons(I),
		}
		Track[I] = TrackCell{
			CurPosR: RevCandidate.CurPosR,
			CurPosC: RevCandidate.CurPosC,
			PrevMov: NewPrev,
			Visited: true,
			Image:   FirstCulImage,
			Start:   RevCandidate.Start,
			Cul:     true,
			Rev:     false,
		}
		fmt.Printf("Previous swapped to: %v\n", NewPrev)

		PickNext(numRows, numCols, I)
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
			Cul:     reverser.Cul,
			Rev:     true,
		}

		/*rev1 := DetermineTrackIconToSet(I)
		rev1ImageSet := IconSet{
			Ic1: rev1,
		}

		Track[I+1] = TrackCell{
			CurPosR: reverser.CurPosR,
			CurPosC: reverser.CurPosC,
			PrevMov: NPrev,
			Visited: reverser.Visited,
			Image:   rev1ImageSet,
			Start:   reverser.Start,
			Cul:     reverser.Cul,
			Rev:     true,
		}*/

		cell := TrackCell{
			CurPosR: -1,
			CurPosC: -1,
			PrevMov: "",
			Visited: false,
			Start:   false,
		}
		Track = append(Track, cell)

		fmt.Printf("Previous swapped to: %v\n", NPrev)

		RevCount++
		fmt.Printf("RevCount: %v\n", RevCount)

		RevInt--
		PickNext(numRows, numCols, I+1)
		return
	}

}
