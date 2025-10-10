package main

import (
	"fmt"
)

var RevInt int

func ReverseProtocol(stop <-chan struct{}, numRows, numCols, I int) bool {

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
			Visited:  1,
			Image:    RevCandidate.Image,
			Start:    RevCandidate.Start,
			Cul:      true,
			Rev:      false,
		}
		fmt.Printf("Previous swapped to: %v\n", NewPrev)

		return PickNext(stop, numRows, numCols, I)
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
			Visited:  reverser.Visited + 1,
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
			Visited:  0,
			Start:    false,
		}
		Track = append(Track, cell)

		fmt.Printf("Previous swapped to: %v\n", NPrev)

		RevCount++
		fmt.Printf("RevCount: %v\n", RevCount)

		RevInt--
		return PickNext(stop, numRows, numCols, I+1)

	}
	fmt.Printf("reverse protocol fail")
	return false
}
