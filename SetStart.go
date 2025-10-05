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

	PMSS := PossibleMovesSS(CurrentRow, CurrentCol, numRows, numCols)

	if len(PMSS) == 0 {
		fmt.Println("Possible moves not being added to list")
	}

	randomIndex := rand.Intn(len(PMSS))

	nextMove := PMSS[randomIndex]

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

	StarticonToSet := DetermineStartImage(1)
	SIconSet := IconSet{Ic1: StarticonToSet}

	Track[0] = TrackCell{
		CurPosR: Track[0].CurPosR,
		CurPosC: Track[0].CurPosC,
		Image:   SIconSet,
		Visited: Track[0].Visited,
		Start:   Track[0].Start,
	}

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
