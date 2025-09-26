package main

import "fyne.io/fyne/v2/canvas"

var ImagePaths = make(map[string]*canvas.Image)

var RallyLogo = []string{"Rally_logo", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rally_mapper_logo.png"}

var StartRight = []string{"StartRight", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startRight.png"}
var StartAngleUR = []string{"StartAngleUR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startAngleUR.png"}
var StartAngleUL = []string{"StartAngleUL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startAngleUL.png"}
var StartUP = []string{"StartUp", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startUp.png"}
var StartLeft = []string{"StartLeft", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startLeft.png"}

var StraightUP = []string{"StraightUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightUP.png"}
var StraightDown = []string{"StraightDown", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDown.png"}
var StraightRight = []string{"StrightRight", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightRight.png"}
var StraightLeft = []string{"StraightLeft", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightLeft.png"}

var StraightDTR = []string{"StraightDTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDTR.png"}
var StraightDTL = []string{"StraightDTL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDTL.png"}
var StraightDBL = []string{"StraightDBL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDBL.png"}
var StraightDBR = []string{"StraightDBR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDBR.png"}

var CurveT_DBL = []string{"CurveT_DBL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveT_DBL.png"}
var CurveT_DBR = []string{"CurveT_DBR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveT_DBR.png"}
var CurveB_DTR = []string{"CurveB_DTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveB_DTR.png"}
var CurveB_DTL = []string{"CurveB_DTL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveB_DTL.png"}
var CurveL_DTR = []string{"CurveL_DTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveL_DTR.png"}
var CurveL_DBR = []string{"CurveL_DBR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveL_DBR.png"}
var CurveR_DTL = []string{"CurveR_DTL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveR_DTL.png"}
var CurveR_DBL = []string{"CurveR_DBL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveR_DBL.png"}

var Curve90B_R = []string{"Curve90B_R", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Curve90B_R.png"}

var Cul_reversalUP = []string{"Cul_reversalUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Cul_reverseUP.png"}
var Cul_reversalDBL = []string{"Cul_reversalDBL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Cul_reversalDBL.png"}

var TwoWayStraightUP = []string{"TwoWayStraightUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/2wayStraightUP.png"}
var TwoWayStraightDTR = []string{"TwoWayStraightDTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/2wayStraightDTR.png"}

//DTR: diagonal Top Right
//DBL: diagonal Bottom Left

/*bad news, it is far easier to flip and rotate the images ouside of using fyne or go, best to just import them here and create a massive switch case.*/
