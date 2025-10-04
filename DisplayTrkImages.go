package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages() {
	var icons IconSet

	go func() {
		StarticonToSet := DetermineStartImage(1)

		IconStart := IconSet{
			Ic1: StarticonToSet,
		}

		SetImageInCell(Track[0].CurPosR, Track[0].CurPosC, IconStart)

		time.Sleep(500 * time.Millisecond)

		for i := 0; i < len(Track); i++ {
			m := Track[i]
			if m.CurPosR != -1 && m.CurPosC != -1 && !m.Start {

				if !m.Cul && !m.Rev && !m.Finish {

					icons = IconSet{
						Ic1: m.Image,
						// Ic2, Ic3, Ic4 can be nil
					}
					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				}

				if m.Cul || m.Rev {
					icons = IconSet{}

					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				}
				if m.Finish {
					finishline, lastIcon := DetermineLastAndFinishIcon(i)
					icons = IconSet{
						Ic1: lastIcon,
						Ic2: finishline,
					}
					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				}

			}
			time.Sleep(50 * time.Millisecond) // Wait between each image
		}

		for i := range Track {
			fmt.Printf("%+v\n", Track[i])
		}
		fmt.Println("all images set")

	}()
}

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
	SIcon := Track[0].Image
	finishPrev := Track[finishInt].PrevMov

	if SIcon == StartUPicon {
		if Track[finishInt].CurPosR > Track[0].CurPosR {
			switch finishPrev {
			case "DUPLeft":
				return finishlineDicon, UnsetPlaceholdericon //unset needs to be cul
			case "Left":
				return finishlineDicon, CurveR_DBLicon
			case "DDownLeft":
				return finishlineDicon, StraightDBLicon
			case "Down":
				return finishlineDicon, CurveT_DBLicon

			}
		}
		if Track[finishInt].CurPosR == Track[0].CurPosR {
			switch finishPrev {
			case "Left":
				return finishlineLicon, StraightLefticon
			case "DDownLeft":
				return finishlineLicon, CurveDTR_Licon
			}
		}
	}

	if SIcon == StartAngleURicon {
		if Track[finishInt].CurPosC == Track[0].CurPosC {
			switch finishPrev {
			case "Down":
				return finishlineDicon, StraightDownicon
			case "DDownLeft":
				return finishlineDicon, CurveDTR_Bicon
			}
		}

		if Track[finishInt].CurPosR > Track[0].CurPosR {
			switch finishPrev {
			case "DDownLeft":
				return finishlineLicon, CurveDTR_Licon
			case "Left":
				return finishlineLicon, StraightLefticon
			}
		}
	}

	if SIcon == StartRighticon {
		if Track[finishInt].CurPosC > Track[0].CurPosC {
			switch finishPrev {
			case "DDownRight":
				return finishlineDicon, UnsetPlaceholdericon //unset needs to be cul
			case "Left":
				return finishlineDicon, CurveR_DBLicon
			case "DDownLeft":
				return finishlineDicon, StraightDBLicon
			case "Down":
				return finishlineDicon, CurveT_DBLicon

			}
		}
		if Track[finishInt].CurPosR > Track[0].CurPosR {
			switch finishPrev {
			case "DDownLeft":
				return finishlineLicon, CurveDTR_Bicon
			case "Down":
				return finishlineLicon, StraightDownicon
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
