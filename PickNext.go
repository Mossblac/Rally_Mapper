package main

import (
	"fmt"
	"math/rand"
)

func PickNext(stop <-chan struct{}, numRows, numCols, I int) bool {
	select {
	case <-stop:
		return false
	default:

		if I+1 == len(Track) && !FindFinish(I) {
			ResetAndTryAgain()

			fmt.Println("All spaces filled, No Finish found, re-run")
			return false
		}

		if RevCount > len(Track)/4 {
			ResetAndTryAgain()

			fmt.Println("Too many reverse, re-run")
			return false
		}

		var PossibleMoves []map[string][]int
		if len(PossibleMoves) != 0 {
			fmt.Printf("Possible moves not zero at PickNext, length: %v", len(PossibleMoves))
			return false
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

		if len(PossibleMoves) == 0 { //if both cul and rev reverse Protocol
			fmt.Println("ReverseProtocol engaged")
			return ReverseProtocol(stop, numRows, numCols, I)
		}

		randomIndex := rand.Intn(len(PossibleMoves))

		nextMove := PossibleMoves[randomIndex]

		var MoveKey string

		for key := range nextMove {
			MoveKey = key
		}

		NPosition := nextMove[MoveKey]

		Track[I+1] = TrackCell{
			TrackInt: I + 1,
			CurPosR:  NPosition[0],
			CurPosC:  NPosition[1],
			PrevMov:  MoveKey,
			Visited:  1,
		}
		if FindFinish(I) {
			TrackLength = I + 1
			TFinish = I + 1
			fin, ficon := DetermineLastAndFinishIcon(I + 1)
			ficonSet := IconSet{Ic1: ficon, Ic2: fin}
			Track[I+1].Image = ficonSet
			Track[I+1].Finish = true

			if ReverseCheck() {
				ResetAndTryAgain()
				return false
			}

			if TrackLength < (numRows*numCols)/3 {
				ResetAndTryAgain()
				return false
			} else {
				for i := 0; i < I+1; i++ {
					var ticon IconSet
					if !Track[i].Start && !Track[i].Finish {
						if !Track[i].Cul && !Track[i].Rev {
							iconToset := DetermineTrackIconToSet(i)
							ticon = IconSet{Ic1: iconToset}
						}
						if Track[i].Cul && !Track[i].Rev {
							PreSetCuls = append(PreSetCuls, Track[i].TrackInt)
							ticon = DetermineCul(i)
						}

						if Track[i].Rev && !Track[i].Cul {
							ticon = DetermineRev(i)
						}

						if Track[i].Cul && Track[i].Rev {
							for _, cul := range PreSetCuls {
								if Track[cul].CurPosR == Track[i].CurPosR && Track[cul].CurPosC == Track[i].CurPosC {
									ticon = Track[cul].Image
								}
							}
						}
						Track[i].Image = ticon
					}
					DetermineCorners(numRows, numCols)
				}
				CalcSpacing()
				DetermineObstacles()
			}
			for i := 0; i < I+1; i++ {
				ic1 := Track[i].Image.Ic1
				ic2 := Track[i].Image.Ic2
				if ic1 == UnsetPlaceholdericon || ic2 == UnsetPlaceholdericon {
					ResetAndTryAgain()
					return false
				}

			}
			fmt.Printf("Found Finish, Track completed\n\n")
			fmt.Printf("TrackLength: %v", TrackLength)
			return true
		}
		if !FindFinish(I) && I+1 < len(Track) {
			PossibleMoves = nil
			return PickNext(stop, numRows, numCols, I+1)
		}
		fmt.Printf("Pick Next Skipped!!!\n\n")
		return false
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
