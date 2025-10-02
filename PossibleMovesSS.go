package main

func PossibleMovesSS(CurrentRow, CurrentCol, numRows, numCols int) (PMSS []map[string][]int) {
	var PossibleMovesSS []map[string][]int
	//UP
	if CurrentRow-1 >= 0 {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMovesSS = append(PossibleMovesSS, UP)
	}

	//DUPRight
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMovesSS = append(PossibleMovesSS, DUPRight)
	}

	//DUPleft
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 && !TrackT {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMovesSS = append(PossibleMovesSS, DUPLeft)
	}

	//Right - last option
	if CurrentCol+1 < numCols {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMovesSS = append(PossibleMovesSS, Right)
	}

	//Left - last option
	if CurrentCol-1 >= 0 && !TrackT {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMovesSS = append(PossibleMovesSS, Left)
	}

	return PossibleMovesSS
}
