package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DetermineTrackIconToSet(I int) (icon *fyne.StaticResource) {
	enter := Track[I].PrevMov
	exit := Track[I+1].PrevMov
	celldirect := fmt.Sprintf("%v, %v", enter, exit)

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
	return nil, nil
}

func DetermineCulAndRevIcons() {
	/*
			CulorRevToSeT := DetermineCullorRevTrackIconToSet(i)

				Track[i] = TrackCell{
				CurPosR: Track[i].CurPosR,
				CurPosC: Track[i].CurPosR,
				PrevMov: Track[i].PrevMov,
				Visited: Track[i].Visited,
				Image:   CulorRevToSet,
				Start:   Track[i].Start,
				Cul:     Track[i].Cul,
				Rev:     Track[i].Rev,
			}
	}*/
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
