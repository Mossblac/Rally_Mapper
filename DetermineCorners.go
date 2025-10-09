package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DetermineCorners(numRows, numCols, I int) {
	var ic3 *fyne.StaticResource

	var CornerTL_list []string
	var CornerTR_list []string
	var CornerBL_list []string
	var CornerBR_list []string

	if Track[I].Image.Ic1 == StartAngleURicon || //DUPRight
		Track[I].Image.Ic1 == StraightDTRicon ||
		Track[I].Image.Ic1 == StraightDBLicon ||
		Track[I].Image.Ic1 == CurveB_DTRicon ||
		Track[I].Image.Ic1 == CurveDTR_Bicon ||
		Track[I].Image.Ic1 == CurveDTR_Licon ||
		Track[I].Image.Ic1 == Cul_DTRicon {

		RowRight := Track[I].CurPosR
		ColRight := Track[I].CurPosC + 1
		RightCell := fmt.Sprintf("%v, %v", RowRight, ColRight)
		CornerTL_list = append(CornerTL_list, RightCell)

		RowUP := Track[I].CurPosR - 1
		ColUP := Track[I].CurPosC
		UPCell := fmt.Sprintf("%v, %v", RowUP, ColUP)
		CornerBR_list = append(CornerBR_list, UPCell)
	}

	if Track[I].Image.Ic1 == StraightDTLicon || //DUPleft
		Track[I].Image.Ic1 == StraightDBRicon ||
		Track[I].Image.Ic1 == CurveDTL_Bicon ||
		Track[I].Image.Ic1 == CurveDTL_Ricon ||
		Track[I].Image.Ic1 == CurveB_DTLicon ||
		Track[I].Image.Ic1 == Cul_DTLicon {

		RowLeft := Track[I].CurPosR
		ColLeft := Track[I].CurPosC - 1
		LeftCell := fmt.Sprintf("%v, %v", RowLeft, ColLeft)
		CornerTR_list = append(CornerTR_list, LeftCell)

		RowUP := Track[I].CurPosR - 1
		ColUP := Track[I].CurPosC
		UPCell := fmt.Sprintf("%v, %v", RowUP, ColUP)
		CornerBL_list = append(CornerBL_list, UPCell)
	}

	//DDownRight
	//DDownleft

	for i := 0; i < len(Track); i++ {
		//this checks if any current track needs a corner, need to figure out how to determine if on a list but not in track.
		Row := Track[i].CurPosR
		Col := Track[i].CurPosC
		CurrentCell := fmt.Sprintf("%v, %v", Row, Col)
		for _, strCTL := range CornerTL_list {
			for _, strBR := range CornerBR_list {
				if CurrentCell == strCTL && CurrentCell != strBR {
					ic3 = CornerTLicon
				}
				if CurrentCell == strBR && CurrentCell != strCTL {
					ic3 = CornerBRicon
				}
				if CurrentCell == strCTL && CurrentCell == strBR {
					ic3 = CornersTL_BRicon
				}
			}
		}
		for _, strCTR := range CornerTR_list {
			for _, strBL := range CornerBL_list {
				if CurrentCell == strCTR && CurrentCell != strBL {
					ic3 = CornerTRicon
				}
				if CurrentCell == strBL && CurrentCell != strCTR {
					ic3 = CornerBLicon
				}
				if CurrentCell == strCTR && CurrentCell == strBL {
					ic3 = CornersTR_BLicon
				}
			}
		}

		CIcons := IconSet{
			Ic1: Track[i].Image.Ic1,
			Ic2: Track[i].Image.Ic2,
			Ic3: ic3,
		}

		Track[i] = TrackCell{
			TrackInt: Track[i].TrackInt,
			CurPosR:  Track[i].CurPosR,
			CurPosC:  Track[i].CurPosC,
			PrevMov:  Track[i].PrevMov,
			Visited:  Track[i].Visited,
			Image:    CIcons,
			Start:    Track[i].Start,
			Finish:   Track[i].Finish,
			Cul:      Track[i].Cul,
			Rev:      Track[i].Rev,
			RevRef:   Track[i].RevRef,
		}

	}
}

/*
str := "123"

	// Convert string to int using strconv.Atoi
	integer, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Failed to convert string to integer: %v", err)
	}
*/
