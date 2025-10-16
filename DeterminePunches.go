package main

import (
	"time"

	"fyne.io/fyne/v2"
)

func DeterminePunches() {
	if len(Obstacles) == 0 {
		return
	}
	PunchImages := []*fyne.StaticResource{
		Punch1icon,
		Punch2icon,
		Punch3icon,
		Punch4icon,
		Punch5icon,
		Punch6icon,
		Punch7icon,
		Punch8icon,
		Punch9icon,
		Punch10icon,
	}

	for i := range TrackLength {
		if Track[i].Image.Ic4 != nil || Track[i].Cul && Track[i+1].Rev && Track[i+2].Rev {
			CulsNeedsPunch := i
			PunchList = append(PunchList, CulsNeedsPunch)
		}
	}
	if len(PunchList) > 10 {
		PunchList = PunchList[:10]
	}

	for i := 0; i < len(PunchList); i++ {
		Track[PunchList[i]].Image.Ic5 = PunchImages[i]
		if i == 0 {
			Track[PunchList[i]].PTime = time.Duration(float64(CalcTime(0, PunchList[i])) * 0.75)

		} else {
			Track[PunchList[i]].PTime = time.Duration(float64(CalcTime(PunchList[i-1], PunchList[i])) * 0.75)
		}

	}
}
