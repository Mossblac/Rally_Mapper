package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
)

func SetStart(CellGrid [][]*fyne.Container, numRows, numCols int) {
	currentMap := Trk[0]
	pos, ok := currentMap["Position"].([]int) //you have to assert type when using interface{}
	if !ok {
		fmt.Println("unable to assert slice of integers")
	}
	CurrentRow := pos[0]
	CurrentCol := pos[1]

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

	FirstMove := map[string]interface{}{
		"Position": []int{NewPosition[0], NewPosition[1]},
		"Previous": MoveKey,
		"TrkIndex": 1,
	}

	Trk = append(Trk, FirstMove)
	TrkInt++

	PossibleMoves = PossibleMoves[:0]

	imageToSet := DetermineStartImage(1)
	start := Trk[0]
	start["Image"] = imageToSet
	SetImageInCell(CellGrid, CurrentRow, CurrentCol, imageToSet)
}

func DetermineStartImage(TrkIndex int) (imageToSet []string) {
	Current := Trk[TrkIndex]
	PrevMove := Current["Previous"]
	switch PrevMove {
	case "UP":
		return StartUP
	case "DUPRight":
		return StartAngleUR
	case "DUPLeft":
		return StartAngleUL
	case "Right":
		return StartRight
	case "Left":
		return StartLeft
	default:
		return nil
	}
}
