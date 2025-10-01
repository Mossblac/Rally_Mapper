package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
)

func SetStart(numRows, numCols int) {
	currentCell := Track[0]

	CurrentRow := currentCell.CurPosR
	CurrentCol := currentCell.CurPosC

	PossibleMovesSS(CurrentRow, CurrentCol, numRows, numCols)

	if len(PossibleMoves) == 0 {
		fmt.Println("Possible moves not being added to list")
	}

	randomIndex := rand.Intn(len(PossibleMoves))

	nextMove := PossibleMoves[randomIndex]

	var MoveKey string

	for key := range nextMove {
		MoveKey = key
	}

	NewPosition := nextMove[MoveKey]

	Track[1] = TrackCell{ //should assign track 1 (after start cell)
		CurPosR: NewPosition[0],
		CurPosC: NewPosition[1],
		PrevMov: MoveKey,
		Visited: true,
	}

	PossibleMoves = PossibleMoves[:0]

	iconToSet := DetermineStartImage(1)

	Track[0] = TrackCell{
		CurPosR: Track[0].CurPosR,
		CurPosC: Track[0].CurPosC,
		Image:   iconToSet,
		Visited: Track[0].Visited,
		Start:   Track[0].Start,
	}

	IconStart := IconSet{
		Ic1: iconToSet,
	}

	SetImageInCell(CurrentRow, CurrentCol, IconStart)

	/*for i := range Track {
		fmt.Printf("%+v\n", Track[i])
	}*/

}

func DetermineStartImage(TrackIndex int) (imageToSet *fyne.StaticResource) {
	Current := Track[TrackIndex]
	PrevMove := Current.PrevMov
	switch PrevMove {
	case "UP":
		return StartUPicon
	case "DUPRight":
		return StartAngleURicon
	case "DUPLeft":
		return StartAngleULicon
	case "Right":
		return StartRighticon
	case "Left":
		return StartLefticon
	default:
		return nil
	}

}
