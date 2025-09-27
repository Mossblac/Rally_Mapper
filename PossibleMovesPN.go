package main

func PossibleMovesPN(CurrentRow, CurrentCol, numRows, numCols, Op, I int) {
	//UP
	vis := VistedCheck(Trk, CurrentRow-1, CurrentCol, I)
	if CurrentRow-1 >= 0 && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 7 || Op == 8) {
		PosMoveUP := []int{CurrentRow - 1, CurrentCol}
		UP := map[string][]int{
			"UP": PosMoveUP,
		}
		PossibleMoves = append(PossibleMoves, UP)
	}

	//DUPRight
	vis = VistedCheck(Trk, CurrentRow-1, CurrentCol+1, I)
	if CurrentRow-1 >= 0 && CurrentCol+1 < numCols && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 7) {
		PosMoveDUR := []int{CurrentRow - 1, CurrentCol + 1}
		DUPRight := map[string][]int{
			"DUPRight": PosMoveDUR,
		}
		PossibleMoves = append(PossibleMoves, DUPRight)
	}

	//DUPleft -
	vis = VistedCheck(Trk, CurrentRow-1, CurrentCol-1, I)
	if CurrentRow-1 >= 0 && CurrentCol-1 >= 0 && !vis && (Op == 1 || Op == 2 || Op == 3 || Op == 8) {
		PosMoveDUL := []int{CurrentRow - 1, CurrentCol - 1}
		DUPLeft := map[string][]int{
			"DUPLeft": PosMoveDUL,
		}
		PossibleMoves = append(PossibleMoves, DUPLeft)
	}

	//Down
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol, I)
	if CurrentRow+1 < numRows && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 7 || Op == 8) {
		PosMoveDown := []int{CurrentRow + 1, CurrentCol}
		Down := map[string][]int{
			"Down": PosMoveDown,
		}
		PossibleMoves = append(PossibleMoves, Down)
	}

	//DDownRight
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol+1, I)
	if CurrentRow+1 < numRows && CurrentCol+1 < numCols && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 7) {
		PosMoveDDR := []int{CurrentRow + 1, CurrentCol + 1}
		DDownRight := map[string][]int{
			"DDownRight": PosMoveDDR,
		}
		PossibleMoves = append(PossibleMoves, DDownRight)
	}

	//DDownLeft
	vis = VistedCheck(Trk, CurrentRow+1, CurrentCol-1, I)
	if CurrentRow+1 < numRows && CurrentCol-1 >= 0 && !vis && (Op == 4 || Op == 5 || Op == 6 || Op == 8) {
		PosMoveDDL := []int{CurrentRow + 1, CurrentCol - 1}
		DDownLeft := map[string][]int{
			"DDownLeft": PosMoveDDL,
		}
		PossibleMoves = append(PossibleMoves, DDownLeft)
	}

	//Right
	vis = VistedCheck(Trk, CurrentRow, CurrentCol+1, I)
	if CurrentCol+1 < numCols && !vis && (Op == 1 || Op == 6 || Op == 7) {
		PosMoveR := []int{CurrentRow, CurrentCol + 1}
		Right := map[string][]int{
			"Right": PosMoveR,
		}
		PossibleMoves = append(PossibleMoves, Right)
	}

	//Left
	vis = VistedCheck(Trk, CurrentRow, CurrentCol-1, I)
	if CurrentCol-1 >= 0 && !vis && (Op == 2 || Op == 5 || Op == 8) {
		PosMoveL := []int{CurrentRow, CurrentCol - 1}
		Left := map[string][]int{
			"Left": PosMoveL,
		}
		PossibleMoves = append(PossibleMoves, Left)
	}

	if len(PossibleMoves) == 0 {
		//Right - last option
		vis = VistedCheck(Trk, CurrentRow, CurrentCol+1, I)
		if CurrentCol+1 < numCols && !vis && (Op == 3 || Op == 4) {
			PosMoveR := []int{CurrentRow, CurrentCol + 1}
			Right := map[string][]int{
				"Right": PosMoveR,
			}
			PossibleMoves = append(PossibleMoves, Right)
		}

		//Left - last option
		vis = VistedCheck(Trk, CurrentRow, CurrentCol-1, I)
		if CurrentCol-1 >= 0 && !vis && (Op == 3 || Op == 4) {
			PosMoveL := []int{CurrentRow, CurrentCol - 1}
			Left := map[string][]int{
				"Left": PosMoveL,
			}
			PossibleMoves = append(PossibleMoves, Left)
		}
	}
}
