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
		var NewPrev = Reversed(prev)
		fmt.Printf("Original Previous value: %v\n", prev)

		Track[I] = TrackCell{
			TrackInt: I,
			CurPosR:  RevCandidate.CurPosR,
			CurPosC:  RevCandidate.CurPosC,
			PrevMov:  NewPrev,
			Visited:  true,
			Image:    RevCandidate.Image,
			Start:    RevCandidate.Start,
			Cul:      true,
			Rev:      false,
		}
		fmt.Printf("Previous swapped to: %v\n", NewPrev)

		PickNext(numRows, numCols, I)
		return
	}

	if RevCandidate.Cul || RevCandidate.Rev {
		reverser := Track[RevInt]
		fmt.Printf("creating reverser @: %v, %v\n", reverser.CurPosR, reverser.CurPosC)
		prev := reverser.PrevMov
		var NPrev = Reversed(prev)
		fmt.Printf("Original Previous value: %v\n", prev)

		Track[I+1] = TrackCell{
			TrackInt: I + 1,
			CurPosR:  reverser.CurPosR,
			CurPosC:  reverser.CurPosC,
			PrevMov:  NPrev,
			Visited:  reverser.Visited,
			Image:    reverser.Image,
			Start:    reverser.Start,
			Cul:      reverser.Cul,
			Rev:      true,
			RevRef:   RevInt,
		}

		cell := TrackCell{
			TrackInt: -1,
			CurPosR:  -1,
			CurPosC:  -1,
			PrevMov:  "",
			Visited:  false,
			Start:    false,
		}
		Track = append(Track, cell)

		fmt.Printf("Previous swapped to: %v\n", NPrev)

		RevCount++
		fmt.Printf("RevCount: %v\n", RevCount)

		RevInt--
		/*if Track[I+1].Cul && Track[I+1].Rev { //testing this
			ReverseProtocol(numRows, numCols, I+1)
			return
		} else {*/
		PickNext(numRows, numCols, I+1)
		return

	}

}
