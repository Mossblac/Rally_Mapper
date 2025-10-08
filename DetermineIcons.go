package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func DetermineTrackIconToSet(I int) (icon *fyne.StaticResource) {
	var exit string
	enter := Track[I].PrevMov
	if Track[I+1].Cul {
		exit = Reversed(Track[I+1].PrevMov)
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

func CulEnterDeterminer(enter string) (Ic1 *fyne.StaticResource) {
	switch enter {
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

	return nil

}

func DetermineCul(I int) (iconset IconSet) {
	var ic1 *fyne.StaticResource
	var ic2 *fyne.StaticResource
	//prev here is the exit of cul

	if !Track[I+1].Cul { // this is the same if I+1 is rev
		enterDir := Reversed(Track[I].PrevMov)
		ic1 = CulEnterDeterminer(enterDir)
		exitDir := Track[I+1].PrevMov
		ic2 = CulExitDeterminer(exitDir)
	}

	if Track[I+1].Cul {
		enterDir := Reversed(Track[I].PrevMov)
		ic1 = CulEnterDeterminer(enterDir)
		exitDir := Reversed(Track[I+1].PrevMov)
		ic2 = CulExitDeterminer(exitDir)
	}

	if Track[I+1].Rev {
		enterDir := Reversed(Track[I].PrevMov)
		ic1 = CulEnterDeterminer(enterDir)
		ic2 = nil
	}

	doubleCulicons := IconSet{
		Ic1: ic1,
		Ic2: ic2,
	}
	return doubleCulicons
}

func CulExitDeterminer(culDirect string) *fyne.StaticResource {
	switch culDirect {
	case "Left":
		return Cul_Righticon
	case "Right":
		return Cul_Lefticon
	case "UP":
		return Cul_Downicon
	case "Down":
		return Cul_UPicon
	case "DUPLeft":
		return Cul_DTLicon
	case "DUPRight":
		return Cul_DTRicon
	case "DDownRight":
		return Cul_DBRicon
	case "DDownLeft":
		return Cul_DBLicon
	}
	return nil
}

func DetermineRev(I int) {
	// this needs to be set within reverse protocol, otherwise revint wont be correct.
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

	if Track[I-1].Rev && !Track[I+1].Cul && !Track[I+1].Rev {
		enter1 = Track[RevInt].PrevMov
		exit1 = Reversed(Track[I].PrevMov)
		celldirect1 = fmt.Sprintf("%v, %v", enter1, exit1)
		ic1 = CellDirectDeterminer(celldirect1)

		enter2 = Track[I].PrevMov
		exit2 = Track[I+1].PrevMov
		celldirect2 = fmt.Sprintf("%v, %v", enter2, exit2)
		ic2 = CellDirectDeterminer(celldirect2)
	}

	if Track[I-1].Rev && Track[I+1].Rev {
		enter1 = Reversed(Track[I+1].PrevMov)
		exit1 = Reversed(Track[I].PrevMov)
		celldirect1 = fmt.Sprintf("%v, %v", enter1, exit1)
		ic1 = CellDirectDeterminer(celldirect1)

		enter2 = Track[I].PrevMov
		exit2 = Track[I+1].PrevMov
		celldirect2 = fmt.Sprintf("%v, %v", enter2, exit2)
		ic2 = CellDirectDeterminer(celldirect2)
	}

	if Track[I+1].Cul && Track[I+1].Rev {
		ic1 = Track[I+1].Image.Ic1
		ic2 = Track[I+1].Image.Ic2
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
