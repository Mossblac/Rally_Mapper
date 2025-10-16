package main

import (
	"time"

	"fyne.io/fyne/v2"
)

func TotalTrackTime() time.Duration {
	return CalcTime(0, TrackLength)
}

func CalcTime(Pstart, Punch int) time.Duration {
	var Total time.Duration
	for i := Pstart; i < Punch; i++ {
		if Track[i].Start {
			Total += 3 * time.Second
		}
		if IsCurbDropOrParkingBlock(Track[i].Image.Ic4) {
			Total += 1500 * time.Millisecond
		}
		if IsLowBridge(Track[i].Image.Ic4) {
			Total += 600 * time.Millisecond
		}
		if Track[i].Cul {
			Total += 3 * time.Second
		} else {
			Total += 400 * time.Millisecond
		}
	}
	return Total

}

func IsCurbDropOrParkingBlock(icon *fyne.StaticResource) bool {
	switch icon {
	case CurbDrop_Verticalicon, CurbDrop_Horizontalicon, CurbDrop_D_TR_BLorBL_TRicon,
		CurbDrop_D_TL_BRorBR_TLicon, CurbDrop_Curve_T_BRorBR_Ticon, CurbDrop_Curve_T_BLorBL_Ticon,
		CurbDrop_Curve_R_TLotTL_Ricon, CurbDrop_Curve_R_BLotBL_Ricon, CurbDrop_Curve_L_TRorTR_Licon,
		CurbDrop_Curve_L_BRorBR_Licon, CurbDrop_Curve_B_TRorTR_Bicon, CurbDrop_Curve_B_TLorTL_Bicon:
		return true
	case ParkingBlock_Verticalicon, ParkingBlock_Horizontalicon, ParkingBlock_D_TR_BLorBL_TRicon,
		ParkingBlock_D_TL_BRorBR_TLicon, ParkingBlock_Curve_T_BRorBR_Ticon, ParkingBlock_Curve_T_BLorBL_Ticon,
		ParkingBlock_Curve_R_TLorTL_Ricon, ParkingBlock_Curve_R_BLorBL_Ricon, ParkingBlock_Curve_L_TRorTR_Licon,
		ParkingBlock_Curve_L_BRorBR_Licon, ParkingBlock_Curve_B_TRorTR_Bicon, ParkingBlock_Curve_B_TLorTL_Bicon:
		return true
	}
	return false
}

func IsLowBridge(icon *fyne.StaticResource) bool {
	switch icon {
	case LowBridge_Verticalicon, LowBridge_Horizontalicon, LowBridge_D_TR_BLorBL_TRicon,
		LowBridge_D_TL_BRorBR_TLicon, LowBridge_Curve_T_BRorBR_Ticon, LowBridge_Curve_T_BLorBL_Ticon,
		LowBridge_Curve_R_TLorTL_Ricon, LowBridge_Curve_R_BLorBL_Ricon, LowBridge_Curve_L_TRorTR_Licon,
		LowBridge_Curve_L_BRorBR_Licon, LowBridge_Curve_B_TRorTR_Bicon, LowBridge_Curve_B_TLorTL_Bicon:
		return true
	}
	return false
}
