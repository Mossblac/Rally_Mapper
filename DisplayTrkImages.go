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
					} else {
						icons = IconSet{}
					}

					SetImageInCell(m.CurPosR, m.CurPosC, icons)
				})
			}
			time.Sleep(200 * time.Millisecond) // Wait between each image
		}

		fyne.Do(func() {
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
