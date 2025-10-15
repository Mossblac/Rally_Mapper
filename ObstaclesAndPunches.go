package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DetermineObstacles() {
	for i := 0; i < len(SpacingList); i++ {
		if !NotPlacable(SpacingList[i]) {
			DetermineOBIconAndPlace(SpacingList[i], Obstacles[i])
		}
		if NotPlacable(SpacingList[i]) {
			BackSpace := RecursiveBack(SpacingList[i])
			if BackSpace != 0 {
				DetermineOBIconAndPlace(BackSpace, Obstacles[i])
			}
			if BackSpace == 0 {
				ForwardSpace := RecursiveForward(SpacingList[i])
				if ForwardSpace != 0 {
					DetermineOBIconAndPlace(ForwardSpace, Obstacles[i])
				}
				if ForwardSpace == 0 {
					fmt.Printf("No placement for obstacle found")
					ResetAndTryAgain()
				}
			}
		}

	}

}

func RecursiveBack(I int) int {
	if Track[I].Image.Ic4 != nil {
		return 0
	}
	if NotPlacable(I) {
		RecursiveBack(I - 1)
	} else {
		return I
	}
	return 0
}

func RecursiveForward(I int) int { // test
	for i := range SpacingList {
		if I == SpacingList[i] {
			return 0
		}
	}
	if NotPlacable(I) {
		RecursiveForward(I + 1)
	} else {
		return I
	}
	return 0
}

func CalcSpacing() { // call before Determine Obstacles
	if len(Obstacles) == 0 {
		return
	} else {
		spacing := TFinish / (len(Obstacles) + 1)
		for i := 1; i < len(Obstacles)+1; i++ {
			SpacingList = append(SpacingList, spacing*i)
		}
		fmt.Printf("%v\n", SpacingList)
		if len(SpacingList) != len(Obstacles) {
			fmt.Printf("error calculating spacing, SpacingList: %v, Obstacles: %v\n", len(SpacingList), len(Obstacles))
		}
	}
}

func NotPlacable(I int) bool {
	icon := Track[I].Image.Ic1
	switch icon {
	case Curve90B_Licon, Curve90L_Bicon, Curve90R_Bicon, Curve90B_Ricon, Curve90L_Ticon, Curve90R_Ticon, Curve90T_Licon, Curve90T_Ricon:
		return true
	}
	if Track[I].Cul {
		return true
	}
	if Track[I].Rev {
		return true
	}
	if Track[I].Start {
		return true
	}
	if Track[I].Finish {
		return true
	}
	return false
}

func DetermineOBIconAndPlace(I int, obstacle string) { // ect, etc...
	var ic4 *fyne.StaticResource
	ic1 := Track[I].Image.Ic1
	if obstacle == "LowBridge" {
		switch ic1 {
		case StraightDownicon, StraightUPicon: //vertical
			ic4 = LowBridge_Verticalicon
		case StraightDBLicon, StraightDTRicon: //diag TR_BL
			ic4 = LowBridge_D_TR_BLorBL_TRicon
		case StraightDBRicon, StraightDTLicon: //diag BR_TL
			ic4 = LowBridge_D_TL_BRorBR_TLicon
		case StraightLefticon, StraightRighticon: // horizontal
			ic4 = LowBridge_Horizontalicon
		case CurveB_DTLicon, CurveDTL_Bicon: // curve B_TL
			ic4 = LowBridge_Curve_B_TLorTL_Bicon
		case CurveB_DTRicon, CurveDTR_Bicon: // curve B_TR
			ic4 = LowBridge_Curve_B_TRorTR_Bicon
		case CurveL_DBRicon, CurveDBR_Licon: // curve L_BR
			ic4 = LowBridge_Curve_L_BRorBR_Licon
		case CurveL_DTRicon, CurveDTR_Licon: //curve L_TR
			ic4 = LowBridge_Curve_L_TRorTR_Licon
		case CurveR_DBLicon, CurveDBL_Ricon: //curve R_BL
			ic4 = LowBridge_Curve_R_BLorBL_Ricon
		case CurveR_DTLicon, CurveDTL_Ricon: //curve R_TL
			ic4 = LowBridge_Curve_R_TLorTL_Ricon
		case CurveT_DBLicon, CurveDBL_Ticon: //curve T_BL
			ic4 = LowBridge_Curve_T_BLorBL_Ticon
		case CurveT_DBRicon, CurveDBR_Ticon: //curve T_BR
			ic4 = LowBridge_Curve_T_BRorBR_Ticon
		}
	}

	if obstacle == "ParkingBlock" {
		switch ic1 {
		case StraightDownicon, StraightUPicon: //vertical
			ic4 = ParkingBlock_Verticalicon
		case StraightDBLicon, StraightDTRicon: //diag TR_BL
			ic4 = ParkingBlock_D_TR_BLorBL_TRicon
		case StraightDBRicon, StraightDTLicon: //diag BR_TL
			ic4 = ParkingBlock_D_TL_BRorBR_TLicon
		case StraightLefticon, StraightRighticon: // horizontal
			ic4 = ParkingBlock_Horizontalicon
		case CurveB_DTLicon, CurveDTL_Bicon: // curve B_TL
			ic4 = ParkingBlock_Curve_B_TLorTL_Bicon
		case CurveB_DTRicon, CurveDTR_Bicon: // curve B_TR
			ic4 = ParkingBlock_Curve_B_TRorTR_Bicon
		case CurveL_DBRicon, CurveDBR_Licon: // curve L_BR
			ic4 = ParkingBlock_Curve_L_BRorBR_Licon
		case CurveL_DTRicon, CurveDTR_Licon: //curve L_TR
			ic4 = ParkingBlock_Curve_L_TRorTR_Licon
		case CurveR_DBLicon, CurveDBL_Ricon: //curve R_BL
			ic4 = ParkingBlock_Curve_R_BLorBL_Ricon
		case CurveR_DTLicon, CurveDTL_Ricon: //curve R_TL
			ic4 = ParkingBlock_Curve_R_TLorTL_Ricon
		case CurveT_DBLicon, CurveDBL_Ticon: //curve T_BL
			ic4 = ParkingBlock_Curve_T_BLorBL_Ticon
		case CurveT_DBRicon, CurveDBR_Ticon: //curve T_BR
			ic4 = ParkingBlock_Curve_T_BRorBR_Ticon
		}
	}

	if obstacle == "Curb_Drop" {
		switch ic1 {
		case StraightDownicon, StraightUPicon: //vertical
			ic4 = CurbDrop_Verticalicon
		case StraightDBLicon, StraightDTRicon: //diag TR_BL
			ic4 = CurbDrop_D_TR_BLorBL_TRicon
		case StraightDBRicon, StraightDTLicon: //diag BR_TL
			ic4 = CurbDrop_D_TL_BRorBR_TLicon
		case StraightLefticon, StraightRighticon: // horizontal
			ic4 = CurbDrop_Horizontalicon
		case CurveB_DTLicon, CurveDTL_Bicon: // curve B_TL
			ic4 = CurbDrop_Curve_B_TLorTL_Bicon
		case CurveB_DTRicon, CurveDTR_Bicon: // curve B_TR
			ic4 = CurbDrop_Curve_B_TRorTR_Bicon
		case CurveL_DBRicon, CurveDBR_Licon: // curve L_BR
			ic4 = CurbDrop_Curve_L_BRorBR_Licon
		case CurveL_DTRicon, CurveDTR_Licon: //curve L_TR
			ic4 = CurbDrop_Curve_L_TRorTR_Licon
		case CurveR_DBLicon, CurveDBL_Ricon: //curve R_BL
			ic4 = CurbDrop_Curve_R_BLotBL_Ricon
		case CurveR_DTLicon, CurveDTL_Ricon: //curve R_TL
			ic4 = CurbDrop_Curve_R_TLotTL_Ricon
		case CurveT_DBLicon, CurveDBL_Ticon: //curve T_BL
			ic4 = CurbDrop_Curve_T_BLorBL_Ticon
		case CurveT_DBRicon, CurveDBR_Ticon: //curve T_BR
			ic4 = CurbDrop_Curve_T_BRorBR_Ticon
		}
	}

	Track[I].Image.Ic4 = ic4

	for i := range TFinish {
		if Track[i].CurPosR == Track[I].CurPosR && Track[i].CurPosC == Track[I].CurPosC {
			Track[i].Image.Ic4 = ic4
		}
	}
}
