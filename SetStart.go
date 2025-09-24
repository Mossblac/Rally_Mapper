package main

import (
	"fmt"
	"math/rand"
)

func SetStart(numRows, numCols int) {
	currentMap := Trk[0]
	pos, ok := currentMap["Position"].([]int)
	if !ok {
		fmt.Println("unable to assert slice of integers")
	}
	CurrentRow := pos[0]
	CurrentCol := pos[1]

	//UP
	if CurrentRow-1 >= 0 {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		PossibleMoves = append(PossibleMoves, PosMoveUP)
	}

	//DUPRight
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		PossibleMoves = append(PossibleMoves, PosMoveDUR)
	}

	//DUPleft
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		PossibleMoves = append(PossibleMoves, PosMoveDUL)
	}

	//Right - last option
	if CurrentCol+1 < numCols {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		PossibleMoves = append(PossibleMoves, PosMoveR)
	}

	//Left - last option
	if CurrentCol-1 >= 0 {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		PossibleMoves = append(PossibleMoves, PosMoveL)
	}

	if len(PossibleMoves) == 0 {
		fmt.Println("Possible moves not being added to list")
	}

	randomIndex := rand.Intn(len(PossibleMoves))

	nextMove := PossibleMoves[randomIndex]

	fmt.Printf("%v", nextMove)

}
