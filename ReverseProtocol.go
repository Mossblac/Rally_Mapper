package main

import "fmt"

func ReverseProtocol(numRows, numCols, I int) {
	RevCandidate := Track[I]
	if !RevCandidate.Cul || !RevCandidate.Rev {
		fmt.Printf("creating Cul @: %v, %v\n", RevCandidate.CurPosR, RevCandidate.CurPosC)
		prev := RevCandidate.PrevMov
		var NewPrev = ""
		fmt.Printf("Original Previous value: %v", prev)
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
		fmt.Printf("Previous swapped to: %v", NewPrev)

		PickNext(numRows, numCols, I)

	}

	if RevCandidate.Cul || RevCandidate.Rev {
		reverser := Track[RevInt]
		fmt.Printf("creating reverser @: %v, %v\n", reverser.CurPosR, reverser.CurPosC)
		prev := reverser.PrevMov
		var NewPrev = ""
		fmt.Printf("Original Previous value: %v", prev)
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

		Track[I+1] = TrackCell{
			CurPosR: reverser.CurPosR,
			CurPosC: reverser.CurPosC,
			PrevMov: NewPrev,
			Visited: true,
			Image:   nil,
			Start:   false,
			Cul:     false,
			Rev:     true,
		}

		RevCount++

		PickNext(numRows, numCols, I+1)
	}

}
