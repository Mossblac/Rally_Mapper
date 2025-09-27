package main

import (
	"fmt"
	"math/rand"
	"time"

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

	PossibleMovesPN(CurrentRow, CurrentCol, numRows, numCols, Op, I)

	if len(Trk) < ((numRows*numCols)-RevCount) && RevCount < ((numRows*numCols)/3) && len(PossibleMoves) == 0 {
		ReverseProtocol(CellGrid, numRows, numCols, I)
	} else {
		PrintTrack(Trk)
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
		"TrkIndex": TrkInt + 1,
	}

	Trk = append(Trk, NextMove)

	fmt.Printf("currentposition: %v %v\n", NPosition[0], NPosition[1])

	if len(Trk) < ((numRows * numCols) - RevCount) {
		go func() {
			PossibleMoves = nil
			fyne.Do(func() {
				TrkInt++
				fmt.Printf("TrkInt: %v\n", TrkInt)
				PickNext(CellGrid, numRows, numCols, TrkInt)
			})
		}()
	} else {
		SetImageInCell(CellGrid, NPosition[0], NPosition[1], RallyLogo, time.Duration(200*(time.Millisecond)))
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

func VistedCheck(Trk []map[string]interface{}, CurrentRow, CurrentCol, I int) bool {
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

func PrintTrack(Trk []map[string]interface{}) {
	for _, Track := range Trk {
		fmt.Printf("%+v\n\n", Track)
	}
}
