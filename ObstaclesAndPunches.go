package main

import "fmt"

func DetermineObstacles(I int) {
	var SpacingList []int

	Spacing := TrackLength / (len(Obstacles) + 1)
	for i := 1; i < len(Obstacles)+1; i++ {
		SpacingList = append(SpacingList, Spacing*i)
	}
	fmt.Printf("%v", SpacingList)
	for _, space := range SpacingList {
		if !Track[space].Cul && !Track[space].Rev && Not90(space) && Track[space].Image.Ic4 == nil {
			//DetermineOBIcon(space)
		} else {
			// DetermineObstacles(space -1)
		}
	}

}

func Not90(I int) bool {
	icon := Track[I].Image.Ic1
	switch icon {
	case Curve90B_Licon, Curve90L_Bicon, Curve90R_Bicon, Curve90B_Ricon, Curve90L_Ticon, Curve90R_Ticon, Curve90T_Licon, Curve90T_Ricon:
		return false
	}
	return true

}

func DetermineOBIcon(I int, obstacle string) { // ect, etc...
	ic1 := Track[I].Image.Ic1
	if obstacle == "LowBridge" {
		switch ic1 {
		case StraightDownicon, StraightUPicon:
			Track[I].Image.Ic4 = LowBridge_Verticalicon
		}
	}
}

/*

L = length of the track (in spaces)

N = number of checkpoints

Then the spacing between each point (including start and end) is:

spacing = L / (N + 1)


Then the position of each checkpoint is:

checkpoint_i = spacing × i     where i = 1 to N

*/
