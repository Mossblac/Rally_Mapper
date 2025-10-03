package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
)

func PickNext(numRows, numCols, I int) {
	if I+1 == len(Track) && !FindFinish(I) {

		ResetAndTryAgain(numRows, numCols)

		fmt.Println("All spaces filled, No Finish found, re-run")
		return
	}

	if RevCount > len(Track)/4 {

		ResetAndTryAgain(numRows, numCols)

		fmt.Println("Too many reverse, re-run")
		return
	}

	var PossibleMoves []map[string][]int
	if len(PossibleMoves) != 0 {
		fmt.Printf("Possible moves not zero at PickNext, length: %v", len(PossibleMoves))
	}

	currentCell := Track[I]

	CurrentRow := currentCell.CurPosR
	CurrentCol := currentCell.CurPosC
	Op := DetermineOptions(I)

	//UP
	vis := VistedCheck(CurrentRow-1, CurrentCol, I)
	if !vis && CurrentRow-1 >= 0 && (Op == 1 || Op == 2 || Op == 3) {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMoves = append(PossibleMoves, UP)
	}

	//DUPRight
	vis = VistedCheck(CurrentRow-1, CurrentCol+1, I)
	if !vis && CurrentRow-1 >= 0 && CurrentCol+1 < numCols && (Op == 1 || Op == 3 || Op == 7) {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMoves = append(PossibleMoves, DUPRight)
	}

	//DUPleft -
	vis = VistedCheck(CurrentRow-1, CurrentCol-1, I)
	if !vis && CurrentRow-1 >= 0 && CurrentCol-1 >= 0 && (Op == 2 || Op == 3 || Op == 8) {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMoves = append(PossibleMoves, DUPLeft)
	}

	//Down
	vis = VistedCheck(CurrentRow+1, CurrentCol, I)
	if !vis && CurrentRow+1 < numRows && (Op == 4 || Op == 5 || Op == 6) {
		PosMoveDown := []int{CurrentRow + 1, CurrentCol}
		Down := map[string][]int{
			"Down": PosMoveDown,
		}
		PossibleMoves = append(PossibleMoves, Down)
	}

	//DDownRight
	vis = VistedCheck(CurrentRow+1, CurrentCol+1, I)
	if !vis && CurrentRow+1 < numRows && CurrentCol+1 < numCols && (Op == 4 || Op == 6 || Op == 7) {
		PosMoveDDR := []int{CurrentRow + 1, CurrentCol + 1}
		DDownRight := map[string][]int{
			"DDownRight": PosMoveDDR,
		}
		PossibleMoves = append(PossibleMoves, DDownRight)
	}

	//DDownLeft
	vis = VistedCheck(CurrentRow+1, CurrentCol-1, I)
	if !vis && CurrentRow+1 < numRows && CurrentCol-1 >= 0 && (Op == 4 || Op == 5 || Op == 8) {
		PosMoveDDL := []int{CurrentRow + 1, CurrentCol - 1}
		DDownLeft := map[string][]int{
			"DDownLeft": PosMoveDDL,
		}
		PossibleMoves = append(PossibleMoves, DDownLeft)
	}

	//Right
	vis = VistedCheck(CurrentRow, CurrentCol+1, I)
	if !vis && CurrentCol+1 < numCols && (Op == 1 || Op == 6 || Op == 7) {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMoves = append(PossibleMoves, Right)
	}

	//Left
	vis = VistedCheck(CurrentRow, CurrentCol-1, I)
	if !vis && CurrentCol-1 >= 0 && (Op == 2 || Op == 5 || Op == 8) {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMoves = append(PossibleMoves, Left)
	}

	if len(PossibleMoves) == 0 {
		//Right - last option
		vis = VistedCheck(CurrentRow, CurrentCol+1, I)
		if !vis && CurrentCol+1 < numCols && (Op == 3 || Op == 4) {
			PosMoveR := []int{CurrentRow, CurrentCol + 1}
			Right := map[string][]int{
				"Right": PosMoveR,
			}
			PossibleMoves = append(PossibleMoves, Right)
		}

		//Left - last option
		vis = VistedCheck(CurrentRow, CurrentCol-1, I)
		if !vis && CurrentCol-1 >= 0 && (Op == 3 || Op == 4) {
			PosMoveL := []int{CurrentRow, CurrentCol - 1}
			Left := map[string][]int{
				"Left": PosMoveL,
			}
			PossibleMoves = append(PossibleMoves, Left)
		}

		//Up - last option
		vis := VistedCheck(CurrentRow-1, CurrentCol, I)
		if !vis && CurrentRow-1 >= 0 && (Op == 7 || Op == 8) {
			PosMoveUP := []int{CurrentRow - 1, CurrentCol}
			UP := map[string][]int{
				"UP": PosMoveUP,
			}
			PossibleMoves = append(PossibleMoves, UP)
		}

		//Down - last option
		vis = VistedCheck(CurrentRow+1, CurrentCol, I)
		if !vis && CurrentRow+1 < numRows && (Op == 7 || Op == 8) {
			PosMoveDown := []int{CurrentRow + 1, CurrentCol}
			Down := map[string][]int{
				"Down": PosMoveDown,
			}
			PossibleMoves = append(PossibleMoves, Down)
		}
	}

	if len(PossibleMoves) == 0 {
		fmt.Println("ReverseProtocol engaged")
		ReverseProtocol(numRows, numCols, I)
		return
	}

	randomIndex := rand.Intn(len(PossibleMoves))

	nextMove := PossibleMoves[randomIndex]

	var MoveKey string

	for key := range nextMove {
		MoveKey = key
	}

	NPosition := nextMove[MoveKey]

	Track[I+1] = TrackCell{
		CurPosR: NPosition[0],
		CurPosC: NPosition[1],
		PrevMov: MoveKey,
		Visited: true,
	}

	if !FindFinish(I) && I+1 < len(Track) {
		go func() {
			PossibleMoves = PossibleMoves[:0]
			fyne.Do(func() {
				PickNext(numRows, numCols, I+1)
			})
		}()
	}
	if FindFinish(I) {
		//SetImageInCell(NPosition[0], NPosition[1], RallyLogo) - set finish line on top of start

		fmt.Println("Found Finish, Track completed")
		return
	}

}

func DetermineOptions(I int) (option int) {
	current := Track[I]
	Prev := current.PrevMov
	switch Prev {
	case "DUPRight":
		return 1 // Up, DUpR, R
	case "DUPLeft":
		return 2 //L, DUpL, Up
	case "UP":
		return 3 //DUPL, UP, DUPR -- (L, R)
	case "Down":
		return 4 //DDL, DDR, Down -- (L, R)
	case "DDownLeft":
		return 5 // left, DDL, down
	case "DDownRight":
		return 6 // Down, DDR, Right
	case "Right":
		return 7 //Right, DUR, DDR, -- (UP, Down)
	case "Left":
		return 8 //Left, DUL, DDL,-- (UP, Down)
	default:
		fmt.Print("Previous did not match movekey label\n")
		return
	}
}

func VistedCheck(CurrentRow, CurrentCol, I int) bool {
	for i := range I {
		Ctrack := Track[i]
		if CurrentRow == Ctrack.CurPosR && CurrentCol == Ctrack.CurPosC {
			return true
		}
	}
	return false
}

func FindFinish(I int) bool {
	next := Track[I+1]
	if TrackT && I > 3 {
		finish := Track[0]
		if next.CurPosR == finish.CurPosR-1 && next.CurPosC == finish.CurPosC ||
			next.CurPosR == finish.CurPosR-1 && next.CurPosC == finish.CurPosC+1 && next.PrevMov != "DDownRight" && next.PrevMov != "DUPLeft" ||
			next.CurPosR == finish.CurPosR && next.CurPosC == finish.CurPosC+1 {
			return true
		}
	}
	if !TrackT && I > 3 {
		if next.CurPosR == 0 && next.CurPosC == 0 ||
			next.CurPosR == 0 && next.CurPosC == 1 ||
			next.CurPosR == 0 && next.CurPosC == 2 {
			return true
		}
	}
	return false
}
