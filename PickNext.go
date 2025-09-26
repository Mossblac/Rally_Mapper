package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
)

func PickNext(CellGrid [][]*fyne.Container, numRows, numCols, I int) {
	currentMap := Trk[I]
	pos, ok := currentMap["Position"].([]int)
	if !ok {
		fmt.Println("unable to assert slice of integers")
	}

	CurrentRow := pos[0]
	CurrentCol := pos[1]
	Op := DetermineOptions(I)

	//UP
	vis := VistedCheck(Trk, CurrentRow-1, CurrentCol)
	if CurrentRow-1 >= 0 && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 7 || Op == 8) {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMoves = append(PossibleMoves, UP)
	}

	//DUPRight
	vis = VistedCheck(Trk, CurrentRow-1, CurrentCol+1)
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 7) {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMoves = append(PossibleMoves, DUPRight)
	}

	//DUPleft -
	vis = VistedCheck(Trk, CurrentRow-1, CurrentCol-1)
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 8) {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMoves = append(PossibleMoves, DUPLeft)
	}

	//Down
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol)
	if CurrentRow+1 < numRows && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 7 || Op == 8) {
		PosMoveDown := []int{CurrentRow + 1, CurrentCol}
		Down := map[string][]int{
			"Down": PosMoveDown,
		}
		PossibleMoves = append(PossibleMoves, Down)
	}

	//DDownRight
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol+1)
	if CurrentRow+1 < numRows && CurrentCol+1 < numCols && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 7) {
		PosMoveDDR := []int{CurrentRow + 1, CurrentCol + 1}
		DDownRight := map[string][]int{
			"DDownRight": PosMoveDDR,
		}
		PossibleMoves = append(PossibleMoves, DDownRight)
	}

	//DDownLeft
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol-1)
	if CurrentRow+1 < numRows && CurrentCol-1 >= 0 && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 8) {
		PosMoveDDL := []int{CurrentRow + 1, CurrentCol - 1}
		DDownLeft := map[string][]int{
			"DDownLeft": PosMoveDDL,
		}
		PossibleMoves = append(PossibleMoves, DDownLeft)
	}

	//Right
	vis = VistedCheck(Trk, CurrentRow, CurrentCol+1)
	if CurrentCol+1 < numCols && !vis && (Op == 1 || Op == 6 || Op == 7) {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMoves = append(PossibleMoves, Right)
	}

	//Left
	vis = VistedCheck(Trk, CurrentRow, CurrentCol-1)
	if CurrentCol-1 >= 0 && !vis && (Op == 2 || Op == 5 || Op == 8) {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMoves = append(PossibleMoves, Left)
	}

	if len(PossibleMoves) == 0 {
		//Right - last option
		vis = VistedCheck(Trk, CurrentRow, CurrentCol+1)
		if CurrentCol+1 < numCols && !vis && (Op == 3 || Op == 4) {
			PosMoveR := []int{CurrentRow, CurrentCol + 1}
			Right := map[string][]int{
				"Right": PosMoveR,
			}
			PossibleMoves = append(PossibleMoves, Right)
		}

		//Left - last option
		vis = VistedCheck(Trk, CurrentRow, CurrentCol-1)
		if CurrentCol-1 >= 0 && !vis && (Op == 3 || Op == 4) {
			PosMoveL := []int{CurrentRow, CurrentCol - 1}
			Left := map[string][]int{
				"Left": PosMoveL,
			}
			PossibleMoves = append(PossibleMoves, Left)
		}
	}

	if len(Trk) < numRows*numCols && len(PossibleMoves) == 0 {
		SetImageInCell(CellGrid, CurrentRow, CurrentCol, RallyLogo)
		fmt.Println("Ran Out of Possible Moves") // here is where you will run the reverse protocol
		return
	}

	randomIndex := rand.Intn(len(PossibleMoves))

	nextMove := PossibleMoves[randomIndex]

	var MoveKey string

	for key := range nextMove {
		MoveKey = key
	}

	NPosition := nextMove[MoveKey]

	NextMove := map[string]interface{}{
		"Position": []int{NPosition[0], NPosition[1]},
		"Previous": MoveKey,
		"TrkIndex": I + 1,
	}

	Trk = append(Trk, NextMove)

	fmt.Printf("currentposition: %v %v\n", NPosition[0], NPosition[1])

	if len(Trk) < numRows*numCols {
		go func() {
			PossibleMoves = nil
			fyne.Do(func() {
				PickNext(CellGrid, numRows, numCols, I+1)
			})
		}()
	} else {
		SetImageInCell(CellGrid, NPosition[0], NPosition[1], RallyLogo)
		fmt.Printf("%v", Trk)
		fmt.Print("map completed")
		return
	}
}

func DetermineOptions(I int) (option int) {
	current := Trk[I]
	Prev := current["Previous"]
	switch Prev {
	case "DUPRight":
		return 1 //DupL, Up, DUpR, R
	case "DUPLeft":
		return 2 //L, DUpL, Up, DupR
	case "UP":
		return 3 //DUPL, UP, DUPR -- (L, R)
	case "Down":
		return 4 //DDL, DDR, Down -- (L, R)
	case "DDownLeft":
		return 5 // left, DDL, down, DDR
	case "DDownRight":
		return 6 // DDL, Down, DDR, Right
	case "Right":
		return 7 //Right, DUR, DDR, UP, Down
	case "Left":
		return 8 //Left, DUL, DDL, UP, Down
	default:
		fmt.Print("Previous did not match movekey label\n")
		return
	}
}

func VistedCheck(Trk []map[string]interface{}, CurrentRow, CurrentCol int) bool {
	for _, Vcheck := range Trk {
		Vis, ok := Vcheck["Position"].([]int)
		if ok {
			x := Vis[0]
			y := Vis[1]
			visited := fmt.Sprintf("%v%v", x, y)
			PosMove := fmt.Sprintf("%v%v", CurrentRow, CurrentCol)
			if visited == PosMove {
				return true
			}

		}

	}
	return false
}
