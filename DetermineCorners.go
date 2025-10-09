package main

func DetermineCorners(numRows, numCols, I int) {

	if Track[I].Image.Ic1 == StartAngleURicon || //DUPRight
		Track[I].Image.Ic1 == StraightDTRicon ||
		Track[I].Image.Ic1 == StraightDBLicon ||
		Track[I].Image.Ic1 == CurveB_DTRicon ||
		Track[I].Image.Ic1 == CurveL_DTRicon ||
		Track[I].Image.Ic1 == CurveDTR_Bicon ||
		Track[I].Image.Ic1 == CurveDTR_Licon ||
		Track[I].Image.Ic1 == Cul_DTRicon {

	}

	if Track[I].Image.Ic1 == StraightDTLicon || //DUPleft
		Track[I].Image.Ic1 == StraightDBRicon ||
		Track[I].Image.Ic1 == CurveDTL_Bicon ||
		Track[I].Image.Ic1 == CurveDTL_Ricon ||
		Track[I].Image.Ic1 == CurveB_DTLicon ||
		Track[I].Image.Ic1 == CurveR_DTLicon ||
		Track[I].Image.Ic1 == Cul_DTLicon {

	}

	//DDownRight
	//DDownleft

}

/*
go
type Cell struct{ R, C int }
type void struct{}
var v void

tl := map[Cell]void{}
tr := map[Cell]void{}
bl := map[Cell]void{}
br := map[Cell]void{}

helper
inBounds := func(r, c, numRows, numCols int) bool {
if r >= 0 && r < numRows && c >= 0 && c < numCols{
return true
}
return false
}

for i := range Track {
    r, c := Track[i].CurPosR, Track[i].CurPosC
    if drivesCorner(Track[i].Images.Ic1) {
        if inBounds(r, c+1) { tl[Cell{r, c+1}] = v }
        if inBounds(r-1, c) { br[Cell{r-1, c}] = v }
    }
}

// go
for i := range Track {
    cell := Cell{Track[i].CurPosR, Track[i].CurPosC}
    inTL := _, has := tl[cell]; has
    inTR := _, has := tr[cell]; has
    inBL := _, has := bl[cell]; has
    inBR := _, has := br[cell]; has

    switch {
    case inTL && inBR:
        Track[i].Images.Ic3 = CornersTL_BRicon
    case inTL:
        Track[i].Images.Ic3 = CornerTLicon
    case inBR:
        Track[i].Images.Ic3 = CornerBRicon
    default:
        // leave as-is or default
    }
}
*/
