package main

import "fyne.io/fyne/v2"

func DeterminePunches() { //after obstacles are placed, changes SpacingList
	//var EndofRevsList []int
	var PunchList []int
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

	for i := range TFinish {
		if Track[i].Image.Ic4 != nil || Track[i].Cul && Track[i+1].Rev && Track[i+2].Rev {
			CulsNeedsPunch := i
			PunchList = append(PunchList, CulsNeedsPunch)
		}
		/*if Track[i].Rev && !Track[i+1].Rev {
			EndofRevint := i
			EndofRevsList = append(EndofRevsList, EndofRevint)
		}*/
	}
	if len(PunchList) > 10 {
		PunchList = PunchList[:10]
	}

	for i := 0; i < len(PunchList); i++ {
		Track[PunchList[i]].Image.Ic5 = PunchImages[i]
	}

}
