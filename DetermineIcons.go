package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func CulConversionForIconToSet(I int) string {
	CulIcon := Track[I].Image.Ic1
	switch CulIcon {
	case Cul_UPicon:
		return "UP"
	case Cul_Downicon:
		return "Down"
	case Cul_Lefticon:
		return "Left"
	case Cul_Righticon:
		return "Right"
	case Cul_DBLicon:
		return "DUPRight"
	case Cul_DBRicon:
		return "DUPLeft"
	case Cul_DTLicon:
		return "DDownRight"
	case Cul_DTRicon:
		return "DDownLeft"
	}
	return ""
}

func DetermineTrackIconToSet(I int) (icon *fyne.StaticResource) {
	var exit string
	enter := Track[I].PrevMov
	if Track[I+1].Cul {
		CulExit := CulConversionForIconToSet(I + 1)
		exit = CulExit
	} else {
		exit = Track[I+1].PrevMov
	}
	celldirect := fmt.Sprintf("%v, %v", enter, exit)
	return CellDirectDeterminer(celldirect)

}

func CellDirectDeterminer(celldirect string) *fyne.StaticResource {
	switch celldirect {
	case "UP, UP":
		return StraightUPicon
	case "Right, Right":
		return StraightRighticon
	case "Left, Left":
		return StraightLefticon
	case "Down, Down":
		return StraightDownicon
	case "DUPRight, DUPRight":
		return StraightDTRicon
	case "DUPLeft, DUPLeft":
		return StraightDTLicon
	case "DDownRight, DDownRight":
		return StraightDBRicon
	case "DDownLeft, DDownLeft":
		return StraightDBLicon
	case "Down, DDownRight":
		return CurveT_DBRicon
	case "Down, DDownLeft":
		return CurveT_DBLicon
	case "Left, DUPLeft":
		return CurveR_DTLicon
	case "Left, DDownLeft":
		return CurveR_DBLicon
	case "Right, DUPRight":
		return CurveL_DTRicon
	case "Right, DDownRight":
		return CurveL_DBRicon
	case "DDownLeft, Left":
		return CurveDTR_Licon
	case "DDownLeft, Down":
		return CurveDTR_Bicon
	case "DDownRight, Right":
		return CurveDTL_Ricon
	case "DDownRight, Down":
		return CurveDTL_Bicon
	case "DUPLeft, UP":
		return CurveDBR_Ticon
	case "DUPRight, UP":
		return CurveDBL_Ticon
	case "UP, DUPRight":
		return CurveB_DTRicon
	case "UP, DUPLeft":
		return CurveB_DTLicon
	case "DUPLeft, Left":
		return CurveDBR_Licon
	case "DUPRight, Right":
		return CurveDBL_Ricon
	case "Down, Right":
		return Curve90T_Ricon
	case "Down, Left":
		return Curve90T_Licon
	case "Left, UP":
		return Curve90R_Ticon
	case "Left, Down":
		return Curve90R_Bicon
	case "Right, UP":
		return Curve90L_Ticon
	case "Right, Down":
		return Curve90L_Bicon
	case "UP, Right":
		return Curve90B_Ricon
	case "UP, Left":
		return Curve90B_Licon
	default:
		return UnsetPlaceholdericon

	}
}

func DetermineLastAndFinishIcon(finishInt int) (finishicon, trackicon *fyne.StaticResource) {
	//return FinishDDLefticon, StraightDBLicon

	SIcon := Track[0].Image.Ic1
	finishPrev := Track[finishInt].PrevMov
	//loop
	if TrackT {
		//Start Up
		if SIcon == StartUPicon {
			//above start - diag
			if Track[finishInt].CurPosR < Track[0].CurPosR && Track[finishInt].CurPosC > Track[0].CurPosC {
				switch finishPrev {
				case "DUPLeft":
					return FinishDDLefticon, UnsetPlaceholdericon //unset needs to be cul
				case "Left":
					return FinishDDLefticon, CurveR_DBLicon
				case "DDownLeft":
					return FinishDDLefticon, StraightDBLicon
				case "Down":
					return FinishDDLefticon, CurveT_DBLicon
				}
			}
			// right of start
			if Track[finishInt].CurPosR == Track[0].CurPosR && Track[finishInt].CurPosC > Track[0].CurPosC {
				switch finishPrev {
				case "Left":
					return FinishLefticon, StraightLefticon
				case "DDownLeft":
					return FinishLefticon, CurveDTR_Licon
				case "Down":
					return FinishDownicon, StraightDownicon
				}
			}
		}
		// Start diag
		if SIcon == StartAngleURicon {
			// right of start
			if Track[finishInt].CurPosC > Track[0].CurPosC && Track[finishInt].CurPosR == Track[0].CurPosR {
				switch finishPrev {
				case "Down":
					return FinishDDLefticon, StraightDownicon
				case "DDownLeft":
					return FinishLefticon, CurveDTR_Licon
				case "DDownRight":
					return FinishDownicon, CurveDTL_Bicon
				case "Left":
					return FinishLefticon, StraightLefticon
				}
			}
			//above start
			if Track[finishInt].CurPosR < Track[0].CurPosR && Track[finishInt].CurPosC == Track[0].CurPosC {
				switch finishPrev {
				case "DDownLeft":
					return FinishDownicon, CurveDTR_Bicon
				case "Down":
					return FinishDownicon, StraightDownicon
				}
			}
		}
		// start right
		if SIcon == StartRighticon {
			// above diag
			if Track[finishInt].CurPosC > Track[0].CurPosC && Track[finishInt].CurPosR < Track[0].CurPosR {
				switch finishPrev {
				case "DDownRight":
					return FinishDDLefticon, UnsetPlaceholdericon //unset needs to be cul
				case "Left":
					return FinishDDLefticon, CurveR_DBLicon
				case "DDownLeft":
					return FinishDDLefticon, StraightDBLicon
				case "Down":
					return FinishDDLefticon, CurveT_DBLicon
				case "DUPLeft":
					return FinishDDLefticon, UnsetPlaceholdericon // also needs to be cul

				}
			}
			//above
			if Track[finishInt].CurPosR < Track[0].CurPosR && Track[finishInt].CurPosC == Track[0].CurPosC {
				switch finishPrev {
				case "DDownLeft":
					return FinishDownicon, CurveDTR_Bicon
				case "Down":
					return FinishDownicon, StraightDownicon
				case "Left":
					return FinishDownicon, Curve90R_Bicon
				}
			}

		}
	}

	if !TrackT {
		//left
		if Track[finishInt].CurPosR == 0 && Track[finishInt].CurPosC == 0 {
			switch finishPrev {
			case "UP":
				return FinishUPicon, StraightUPicon
			case "DUPLeft":
				return FinishUPicon, CurveDBR_Ticon
			}
		}
		// center
		if Track[finishInt].CurPosR == 0 && Track[finishInt].CurPosC == 1 {
			switch finishPrev {
			case "UP":
				return FinishUPicon, StraightUPicon
			case "DUPLeft":
				return FinishUPicon, CurveDBR_Ticon
			case "DUPRight":
				return FinishUPicon, CurveDBL_Ticon
			}
		}

		//right
		if Track[finishInt].CurPosR == 0 && Track[finishInt].CurPosC == 2 {
			switch finishPrev {
			case "UP":
				return FinishUPicon, StraightUPicon
			case "DUPRight":
				return FinishUPicon, CurveDBL_Ticon
			}
		}

	}
	return nil, nil
}

func DetermineFirstCulIcons(I int) (Ic1 *fyne.StaticResource) {

	preCulpreMove := Track[I].PrevMov //this is triple checked, working
	switch preCulpreMove {
	case "Left":
		return Cul_Lefticon
	case "Right":
		return Cul_Righticon
	case "UP":
		return Cul_UPicon
	case "Down":
		return Cul_Downicon
	case "DUPLeft":
		return Cul_DBRicon
	case "DUPRight":
		return Cul_DBLicon
	case "DDownRight":
		return Cul_DTLicon
	case "DDownLeft":
		return Cul_DTRicon
	}

	/*if !Track[I+1].Cul && !Track[I+1].Rev {
		preCulpreMove := Track[I].PrevMov
		switch preCulpreMove {
		case "Left":
			return Cul_Righticon
		case "Right":
			return Cul_Lefticon
		case "UP":
			return Cul_Downicon
		case "Down":
			return Cul_UPicon
		case "DUPLeft":
			return Cul_DBRicon
		case "DUPRight":
			return Cul_DBLicon
		case "DDownRight":
			return Cul_DTLicon
		case "DDownLeft":
			return Cul_DTRicon
		}
	}*/
	return nil

}

func DetermineSecondCul(I int) (iconset IconSet) {
	//prev here is the exit of cul
	var ic2 *fyne.StaticResource

	if !Track[I+1].Cul && !Track[I+1].Rev { //this is double checked
		preCulpreMove := Track[I+1].PrevMov
		switch preCulpreMove {
		case "Left":
			ic2 = Cul_Righticon
		case "Right":
			ic2 = Cul_Lefticon
		case "UP":
			ic2 = Cul_Downicon
		case "Down":
			ic2 = Cul_UPicon
		case "DUPLeft":
			ic2 = Cul_DTLicon
		case "DUPRight":
			ic2 = Cul_DTRicon
		case "DDownRight":
			ic2 = Cul_DBRicon
		case "DDownLeft":
			ic2 = Cul_DBLicon
		}
	}
	doubleCulicons := IconSet{
		Ic1: Track[I].Image.Ic1,
		Ic2: ic2,
	}
	return doubleCulicons
}

func DetermineRev(I int) {
	var ic1 *fyne.StaticResource
	var ic2 *fyne.StaticResource
	var enter1 string
	var exit1 string       // 1 = towards cul, before reverse
	var celldirect1 string // 2 = away from cull after reverse
	var enter2 string
	var exit2 string
	var celldirect2 string
	if !Track[I+1].Rev && !Track[I+1].Cul {
		//towards cul
		enter1 = Track[RevInt].PrevMov
		exit1 = Reversed(Track[I].PrevMov)
		celldirect1 = fmt.Sprintf("%v, %v", enter1, exit1)
		ic1 = CellDirectDeterminer(celldirect1)
		// away from cul
		enter2 = Track[I].PrevMov
		exit2 = Track[I+1].PrevMov
		celldirect2 = fmt.Sprintf("%v, %v", enter2, exit2)
		ic2 = CellDirectDeterminer(celldirect2)
	}

	RevSet := IconSet{
		Ic1: ic1,
		Ic2: ic2, // set for testing
	}

	Track[I] = TrackCell{
		CurPosR: Track[I].CurPosR,
		CurPosC: Track[I].CurPosC,
		PrevMov: Track[I].PrevMov,
		Visited: Track[I].Visited,
		Image:   RevSet,
		Start:   Track[I].Start,
		Finish:  Track[I].Finish,
		Cul:     Track[I].Cul,
		Rev:     Track[I].Rev,
	}
}

func DetermineCorners(I int) {

}

func Reversed(input string) string {
	input = ""
	switch input {
	case "UP":
		return "Down"
	case "Down":
		return "UP"
	case "Left":
		return "Right"
	case "Right":
		return "Left"
	case "DDownRight":
		return "DUPLeft"
	case "DDownLeft":
		return "DUPRight"
	case "DUPRight":
		return "DDownLeft"
	case "DUPLeft":
		return "DDownRight"
	default:
		return ""
	}
}

/*

UP
DUPRight
DUPLeft
Down
DDownRight
DDownLeft
Right
Left

*/

/*if Track[I+1].Cul && !Track[I+1].Rev {
	preCulpreMove := Track[I+1].PrevMov
	switch preCulpreMove {
	case "Left":
		ic2 = Cul_Lefticon
	case "Right":
		ic2 = Cul_Righticon
	case "UP":
		ic2 = Cul_UPicon
	case "Down":
		ic2 = Cul_Downicon
	case "DUPLeft":
		ic2 = Cul_DTLicon
	case "DUPRight":
		ic2 = Cul_DTRicon
	case "DDownRight":
		ic2 = Cul_DBRicon
	case "DDownLeft":
		ic2 = Cul_DBLicon
	}
}*/

/*var ic2 *fyne.StaticResource
enter := Track[I].PrevMov
exit := Track[I+1].PrevMov
celldirect := fmt.Sprintf("%v, %v", enter, exit)
switch celldirect {
case "UP, UP":
	ic2 = Rev_Verticalicon
case "Down, Down":
	ic2 = Rev_Verticalicon
case "DUPRight, DUPRight":
	ic2 = Rev_Diag_TR_BLicon
case "DDownLeft, DDownLeft":
	ic2 = Rev_Diag_TR_BLicon
case "DUPLeft, DUPLeft":
	ic2 = Rev_Diag_TL_BRicon
case "DDownRight, DDownRight":
	ic2 = Rev_Diag_TL_BRicon
case "Left, Left":
	ic2 = Rev_Horizontalicon
case "Right, Right":
	ic2 = Rev_Horizontalicon
	// it is the curves next....
}*/
