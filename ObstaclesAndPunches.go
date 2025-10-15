package main

import "fmt"

var SpacingList []int

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

func RecursiveForward(I int) int { // needs to be next in Spacing list, not the number plus 1
	if I == SpacingList[I+1] {
		return 0
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
		Spacing := TrackLength / (len(Obstacles) + 1)
		for i := 1; i < len(Obstacles)+1; i++ {
			SpacingList = append(SpacingList, Spacing*i)
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
	Track[I].Image.Ic4 = UnsetPlaceholdericon
	/*ic1 := Track[I].Image.Ic1
	if obstacle == "LowBridge" {
		switch ic1 {
		case StraightDownicon, StraightUPicon:
			Track[I].Image.Ic4 = LowBridge_Verticalicon
		}
	}*/
}
