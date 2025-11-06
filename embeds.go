package main

import _ "embed"

//go:embed assets/icons/StartAngleUL.svg
var StartAngleULSVG []byte

//go:embed assets/icons/StartAngleUR.svg
var StartAngleURSVG []byte

//go:embed assets/icons/StartLeft.svg
var StartLeftSVG []byte

//go:embed assets/icons/StartRight.svg
var StartRightSVG []byte

//go:embed assets/icons/StartUP.svg
var StartUPSVG []byte

//go:embed assets/icons/StraightUP.svg
var StraightUPSVG []byte

//go:embed assets/icons/UnsetPlaceholder.svg
var UnsetPlaceholderSVG []byte

//go:embed assets/icons/StraightDown.svg
var StraightDownSVG []byte

//go:embed assets/icons/StraightLeft.svg
var StraightLeftSVG []byte

//go:embed assets/icons/StraightRight.svg
var StraightRightSVG []byte

//go:embed assets/icons/StraightDBL.svg
var StraightDBLSVG []byte

//go:embed assets/icons/StraightDBR.svg
var StraightDBRSVG []byte

//go:embed assets/icons/StraightDTL.svg
var StraightDTLSVG []byte

//go:embed assets/icons/StraightDTR.svg
var StraightDTRSVG []byte

//go:embed assets/icons/CurveT_DBR.svg
var CurveT_DBRSVG []byte

//go:embed assets/icons/CurveT_DBL.svg
var CurveT_DBLSVG []byte

//go:embed assets/icons/CurveR_DTL.svg
var CurveR_DTLSVG []byte

//go:embed assets/icons/CurveR_DBL.svg
var CurveR_DBLSVG []byte

//go:embed assets/icons/CurveL_DTR.svg
var CurveL_DTRSVG []byte

//go:embed assets/icons/CurveL_DBR.svg
var CurveL_DBRSVG []byte

//go:embed assets/icons/CurveDTR_L.svg
var CurveDTR_LSVG []byte

//go:embed assets/icons/CurveDTR_B.svg
var CurveDTR_BSVG []byte

//go:embed assets/icons/CurveDTL_R.svg
var CurveDTL_RSVG []byte

//go:embed assets/icons/CurveDTL_B.svg
var CurveDTL_BSVG []byte

//go:embed assets/icons/CurveDBR_T.svg
var CurveDBR_TSVG []byte

//go:embed assets/icons/CurveDBL_T.svg
var CurveDBL_TSVG []byte

//go:embed assets/icons/CurveB_DTR.svg
var CurveB_DTRSVG []byte

//go:embed assets/icons/CurveB_DTL.svg
var CurveB_DTLSVG []byte

//go:embed assets/icons/CurveDBL_R.svg
var CurveDBL_RSVG []byte

//go:embed assets/icons/CurveDBR_L.svg
var CurveDBR_LSVG []byte

//go:embed assets/icons/Curve90T_R.svg
var Curve90T_RSVG []byte

//go:embed assets/icons/Curve90T_L.svg
var Curve90T_LSVG []byte

//go:embed assets/icons/Curve90R_T.svg
var Curve90R_TSVG []byte

//go:embed assets/icons/Curve90R_B.svg
var Curve90R_BSVG []byte

//go:embed assets/icons/Curve90L_T.svg
var Curve90L_TSVG []byte

//go:embed assets/icons/Curve90L_B.svg
var Curve90L_BSVG []byte

//go:embed assets/icons/Curve90B_R.svg
var Curve90B_RSVG []byte

//go:embed assets/icons/Curve90B_L.svg
var Curve90B_LSVG []byte

//go:embed assets/icons/FinishlineB.svg
var FinishlineBSVG []byte

//go:embed assets/icons/FinishlineD.svg
var FinishlineDSVG []byte

//go:embed assets/icons/FinishlineL.svg
var FinishlineLSVG []byte

//go:embed assets/icons/FinishDDLeft.svg
var FinishDDLeftSVG []byte

//go:embed assets/icons/FinishDown.svg
var FinishDownSVG []byte

//go:embed assets/icons/FinishLeft.svg
var FinishLeftSVG []byte

//go:embed assets/icons/FinishUP.svg
var FinishUPSVG []byte

//go:embed assets/icons/Cul_Down.svg
var Cul_DownSVG []byte

//go:embed assets/icons/Cul_Left.svg
var Cul_LeftSVG []byte

//go:embed assets/icons/Cul_UP.svg
var Cul_UPSVG []byte

//go:embed assets/icons/Cul_Right.svg
var Cul_RightSVG []byte

//go:embed assets/icons/Cul_DBL.svg
var Cul_DBLSVG []byte

//go:embed assets/icons/Cul_DTL.svg
var Cul_DTLSVG []byte

//go:embed assets/icons/Cul_DTR.svg
var Cul_DTRSVG []byte

//go:embed assets/icons/Cul_DBR.svg
var Cul_DBRSVG []byte

//go:embed assets/icons/CornerBL.svg
var CornerBLSVG []byte

//go:embed assets/icons/CornerBR.svg
var CornerBRSVG []byte

//go:embed assets/icons/CornersTL_BR.svg
var CornersTL_BRSVG []byte

//go:embed assets/icons/CornersTR_BL.svg
var CornersTR_BLSVG []byte

//go:embed assets/icons/CornerTL.svg
var CornerTLSVG []byte

//go:embed assets/icons/CornerTR.svg
var CornerTRSVG []byte

//go:embed assets/icons/Rev_Diag_TL_BR.svg
var Rev_Diag_TL_BRSVG []byte

//go:embed assets/icons/Rev_Diag_TR_BL.svg
var Rev_Diag_TR_BLSVG []byte

//go:embed assets/icons/Rev_Horizontal.svg
var Rev_HorizontalSVG []byte

//go:embed assets/icons/Rev_Vertical.svg
var Rev_VerticalSVG []byte

//go:embed assets/icons/Grass.svg
var GrassSVG []byte

//go:embed assets/icons/CurbDrop_Curve_B_TLorTL_B.svg
var CurbDrop_Curve_B_TLorTL_BSVG []byte

//go:embed assets/icons/CurbDrop_Curve_B_TRorTR_B.svg
var CurbDrop_Curve_B_TRorTR_BSVG []byte

//go:embed assets/icons/CurbDrop_Curve_L_BRorBR_L.svg
var CurbDrop_Curve_L_BRorBR_LSVG []byte

//go:embed assets/icons/CurbDrop_Curve_L_TRorTR_L.svg
var CurbDrop_Curve_L_TRorTR_LSVG []byte

//go:embed assets/icons/CurbDrop_Curve_R_BLotBL_R.svg
var CurbDrop_Curve_R_BLotBL_RSVG []byte

//go:embed assets/icons/CurbDrop_Curve_R_TLotTL_R.svg
var CurbDrop_Curve_R_TLotTL_RSVG []byte

//go:embed assets/icons/CurbDrop_Curve_T_BLorBL_T.svg
var CurbDrop_Curve_T_BLorBL_TSVG []byte

//go:embed assets/icons/CurbDrop_Curve_T_BRorBR_T.svg
var CurbDrop_Curve_T_BRorBR_TSVG []byte

//go:embed assets/icons/CurbDrop_D_TL_BRorBR_TL.svg
var CurbDrop_D_TL_BRorBR_TLSVG []byte

//go:embed assets/icons/CurbDrop_D_TR_BLorBL_TR.svg
var CurbDrop_D_TR_BLorBL_TRSVG []byte

//go:embed assets/icons/CurbDrop_Horizontal.svg
var CurbDrop_HorizontalSVG []byte

//go:embed assets/icons/CurbDrop_Vertical.svg
var CurbDrop_VerticalSVG []byte

//go:embed assets/icons/LowBridge_Curve_B_TLorTL_B.svg
var LowBridge_Curve_B_TLorTL_BSVG []byte

//go:embed assets/icons/LowBridge_Curve_B_TRorTR_B.svg
var LowBridge_Curve_B_TRorTR_BSVG []byte

//go:embed assets/icons/LowBridge_Curve_L_BRorBR_L.svg
var LowBridge_Curve_L_BRorBR_LSVG []byte

//go:embed assets/icons/LowBridge_Curve_L_TRorTR_L.svg
var LowBridge_Curve_L_TRorTR_LSVG []byte

//go:embed assets/icons/LowBridge_Curve_R_BLorBL_R.svg
var LowBridge_Curve_R_BLorBL_RSVG []byte

//go:embed assets/icons/LowBridge_Curve_R_TLorTL_R.svg
var LowBridge_Curve_R_TLorTL_RSVG []byte

//go:embed assets/icons/LowBridge_Curve_T_BLorBL_T.svg
var LowBridge_Curve_T_BLorBL_TSVG []byte

//go:embed assets/icons/LowBridge_Curve_T_BRorBR_T.svg
var LowBridge_Curve_T_BRorBR_TSVG []byte

//go:embed assets/icons/LowBridge_D_TL_BRorBR_TL.svg
var LowBridge_D_TL_BRorBR_TLSVG []byte

//go:embed assets/icons/LowBridge_D_TR_BLorBL_TR.svg
var LowBridge_D_TR_BLorBL_TRSVG []byte

//go:embed assets/icons/LowBridge_Horizontal.svg
var LowBridge_HorizontalSVG []byte

//go:embed assets/icons/LowBridge_Vertical.svg
var LowBridge_VerticalSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_B_TLorTL_B.svg
var ParkingBlock_Curve_B_TLorTL_BSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_B_TRorTR_B.svg
var ParkingBlock_Curve_B_TRorTR_BSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_L_BRorBR_L.svg
var ParkingBlock_Curve_L_BRorBR_LSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_L_TRorTR_L.svg
var ParkingBlock_Curve_L_TRorTR_LSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_R_BLorBL_R.svg
var ParkingBlock_Curve_R_BLorBL_RSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_R_TLorTL_R.svg
var ParkingBlock_Curve_R_TLorTL_RSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_T_BLorBL_T.svg
var ParkingBlock_Curve_T_BLorBL_TSVG []byte

//go:embed assets/icons/ParkingBlock_Curve_T_BRorBR_T.svg
var ParkingBlock_Curve_T_BRorBR_TSVG []byte

//go:embed assets/icons/ParkingBlock_D_TL_BRorBR_TL.svg
var ParkingBlock_D_TL_BRorBR_TLSVG []byte

//go:embed assets/icons/ParkingBlock_D_TR_BLorBL_TR.svg
var ParkingBlock_D_TR_BLorBL_TRSVG []byte

//go:embed assets/icons/ParkingBlock_Horizontal.svg
var ParkingBlock_HorizontalSVG []byte

//go:embed assets/icons/ParkingBlock_Vertical.svg
var ParkingBlock_VerticalSVG []byte

//go:embed assets/icons/Punch10.svg
var Punch10SVG []byte

//go:embed assets/icons/Punch1.svg
var Punch1SVG []byte

//go:embed assets/icons/Punch2.svg
var Punch2SVG []byte

//go:embed assets/icons/Punch3.svg
var Punch3SVG []byte

//go:embed assets/icons/Punch4.svg
var Punch4SVG []byte

//go:embed assets/icons/Punch5.svg
var Punch5SVG []byte

//go:embed assets/icons/Punch6.svg
var Punch6SVG []byte

//go:embed assets/icons/Punch7.svg
var Punch7SVG []byte

//go:embed assets/icons/Punch8.svg
var Punch8SVG []byte

//go:embed assets/icons/Punch9.svg
var Punch9SVG []byte
