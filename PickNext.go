package main

import (
	"fmt"
)

var PossibleMoves []map[string][]int

func PickNext(numRows, numCols, I int) {
	currentMap := Trk[I]
	pos, ok := currentMap["Position"].([]int)
	if !ok {
		fmt.Println("unable to assert slice of integers")
	}
	vis, ok := currentMap["Visited"].(bool)
	if !ok {
		fmt.Println("unable to assert Visited is boolean")
		vis = false
	}

	CurrentRow := pos[0]
	CurrentCol := pos[1]

	//UP
	if CurrentRow-1 >= 0 && !vis {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMoves = append(PossibleMoves, UP)
	}

	//DUPRight
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols && !vis {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMoves = append(PossibleMoves, DUPRight)
	}

	//DUPleft
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 && !vis {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMoves = append(PossibleMoves, DUPLeft)
	}

	//Down
	if CurrentRow+1 < numRows && !vis {
		PosMoveDown := []int{CurrentRow + 1, CurrentCol}
		Down := map[string][]int{
			"Down": PosMoveDown,
		}
		PossibleMoves = append(PossibleMoves, Down)
	}

	//DDownRight
	if CurrentRow+1 < numRows && CurrentCol+1 < numCols && !vis {
		PosMoveDDR := []int{CurrentRow + 1, CurrentCol + 1}
		DDownRight := map[string][]int{
			"DDownRight": PosMoveDDR,
		}
		PossibleMoves = append(PossibleMoves, DDownRight)
	}

	//DDownLeft
	if CurrentRow+1 < numRows && CurrentCol-1 >= 0 && !vis {
		PosMoveDDL := []int{CurrentRow + 1, CurrentCol - 1}
		DDownLeft := map[string][]int{
			"DDownLeft": PosMoveDDL,
		}
		PossibleMoves = append(PossibleMoves, DDownLeft)
	}

	//Right - last option
	if CurrentCol+1 < numCols && !vis {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMoves = append(PossibleMoves, Right)
	}

	//Left - last option
	if CurrentCol-1 >= 0 && !vis {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMoves = append(PossibleMoves, Left)
	}

	/*if len(PossibleMoves) == 0 {
		fmt.Println("Possible moves not being added to list")
	}

	randomIndex := rand.Intn(len(PossibleMoves))

	nextMove := PossibleMoves[randomIndex]

	var MoveKey string

	for key := range nextMove {
		MoveKey = key
	}

	NewPosition := nextMove[MoveKey]

	NextMove := map[string]interface{}{
		"Position": []int{NewPosition[0], NewPosition[1]},
		"Visited":  true,
		"Previous": MoveKey,
		"TrkIndex": I+1,
	}

	Trk = append(Trk, NextMove)*/

}

/*func DetermineStartImage(TrkIndex int) (imageToSet []string) {
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
}*/
