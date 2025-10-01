package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
)

func SetStart(numRows, numCols int) {
	currentCell := Track[0]

	CurrentRow := currentCell.CurPosX
	CurrentCol := currentCell.CurPosY

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
		CurPosX: NewPosition[0],
		CurPosY: NewPosition[1],
		PrevMov: MoveKey,
		Visited: true,
	}

	PossibleMoves = PossibleMoves[:0]

	iconToSet := DetermineStartImage(1)

	Track[0] = TrackCell{
		Image: iconToSet,
	}

	IconStart := IconSet{
		Ic1: iconToSet,
	}

	SetImageInCell(CurrentRow, CurrentCol, IconStart)
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
