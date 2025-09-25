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

	//UP
	if CurrentRow-1 >= 0 {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMoves = append(PossibleMoves, UP)
	}

	//DUPRight
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMoves = append(PossibleMoves, DUPRight)
	}

	//DUPleft
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMoves = append(PossibleMoves, DUPLeft)
	}

	//Right - last option
	if CurrentCol+1 < numCols {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMoves = append(PossibleMoves, Right)
	}

	//Left - last option
	if CurrentCol-1 >= 0 {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMoves = append(PossibleMoves, Left)
	}

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
		"Visited":  true,
		"Previous": MoveKey,
		"TrkIndex": 1,
	}

	Trk = append(Trk, FirstMove)

	fmt.Printf("%v", MoveKey)

	PossibleMoves = nil

	imageToSet := DetermineStartImage(1)
	FirstMove["Image"] = imageToSet
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
