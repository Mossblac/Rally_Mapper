package main

import "fmt"

func PickNext(numRows, numCols, I int) {
	var PossibleMoves [][]int
	currentMap := Trk[I]
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

	//Down
	if CurrentRow+1 < numRows {
		PosMoveDown := []int{CurrentRow + 1, CurrentCol}
		PossibleMoves = append(PossibleMoves, PosMoveDown)
	}

	//DDownRight
	if CurrentRow+1 < numRows && CurrentCol+1 < numCols {
		PosMoveDDR := []int{CurrentRow + 1, CurrentCol + 1}
		PossibleMoves = append(PossibleMoves, PosMoveDDR)
	}

	//DDownLeft
	if CurrentRow+1 < numRows && CurrentCol-1 >= 0 {
		PosMoveDDL := []int{CurrentRow + 1, CurrentCol - 1}
		PossibleMoves = append(PossibleMoves, PosMoveDDL)
	}

}
