package main

import "fyne.io/fyne/v2/canvas"

var ImagePaths = make(map[string]*canvas.Image)

var StartRight = []string{"StartRight", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startRight.png"}
var StartAngleUR = []string{"StartAngleUR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startAngleUR.png"}
var StartAngleUL = []string{"StartAngleUL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startAngleUL.png"}
var StartUP = []string{"StartUp", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startUp.png"}
var StartLeft = []string{"StartLeft", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/startLeft.png"}

var StraightUP = []string{"StraightUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightUP.png"}
var StraightDTR = []string{"StraightDTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/StraightDTR.png"}

var CurveB_DTR = []string{"CurveB_DTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/CurveB_DTR.png"}
var Curve90B_R = []string{"Curve90B_R", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Curve90B_R.png"}

var Cul_reversalUP = []string{"Cul_reversalUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Cul_reverseUP.png"}
var Cul_reversalDBL = []string{"Cul_reversalDBL", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/Cul_reversalDBL.png"}

var TwoWayStraightUP = []string{"TwoWayStraightUP", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/2wayStraightUP.png"}
var TwoWayStraightDTR = []string{"TwoWayStraightDTR", "/home/mossblac/workspace/github.com/Mossblac/Rally_Mapper/images/rallyicons/2wayStraightDTR.png"}

//DTR: diagonal Top Right
//DBL: diagonal Bottom Left

/*bad news, it is far easier to flip and rotate the images ouside of using fyne or go, best to just import them here and create a massive switch case.*/
