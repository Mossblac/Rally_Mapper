package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

func DisplayTrkImages() {
	var icons IconSet

	go func() {
		for i := 0; i < len(Track); i++ {
			m := Track[i]
			if m.CurPosR != -1 && m.CurPosC != -1 && !m.Start {

				fyne.Do(func() {
					if !m.Cul && !m.Rev {
						iconToset := DetermineTrackIconToSet(i)

						icons = IconSet{
							Ic1: iconToset,
							// Ic2, Ic3, Ic4 can be nil
						}

						Track[i] = TrackCell{
							CurPosR: Track[i].CurPosR,
							CurPosC: Track[i].CurPosR,
							PrevMov: Track[i].PrevMov,
							Visited: Track[i].Visited,
							Image:   iconToset,
							Start:   Track[i].Start,
							Cul:     Track[i].Cul,
							Rev:     Track[i].Rev,
						}

					} else {
						icons = IconSet{}
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

						*/
					}

					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				})
			}
			time.Sleep(200 * time.Millisecond) // Wait between each image
		}

		fyne.Do(func() {
			for i := range Track {
				fmt.Printf("%+v\n", Track[i])
			}
			fmt.Println("all images set")
		})
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
	case "Left,Left":
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
	default:
		return UnsetPlaceholdericon

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
