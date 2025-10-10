package main

func DetermineCorners(numRows, numCols int) {
	type Cell struct{ R, C int }
	type void struct{}
	var v void

	tl := map[Cell]void{}
	tr := map[Cell]void{}
	bl := map[Cell]void{}
	br := map[Cell]void{}

	for i := range Track {
		r, c := Track[i].CurPosR, Track[i].CurPosC
		if Track[i].Image.Ic1 == StartAngleURicon || //DUPRight
			Track[i].Image.Ic1 == StraightDTRicon ||
			Track[i].Image.Ic1 == StraightDBLicon ||
			Track[i].Image.Ic1 == CurveB_DTRicon ||
			Track[i].Image.Ic1 == CurveL_DTRicon ||
			Track[i].Image.Ic1 == CurveDTR_Bicon ||
			Track[i].Image.Ic1 == CurveDTR_Licon ||
			Track[i].Image.Ic1 == Cul_DTRicon {
			if InBounds(r, c+1, numRows, numCols) {
				tl[Cell{r, c + 1}] = v
			}
			if InBounds(r-1, c, numRows, numCols) {
				br[Cell{r - 1, c}] = v
			}
		}

		if Track[i].Image.Ic1 == StraightDTLicon || //DUPleft
			Track[i].Image.Ic1 == StraightDBRicon ||
			Track[i].Image.Ic1 == CurveDTL_Bicon ||
			Track[i].Image.Ic1 == CurveDTL_Ricon ||
			Track[i].Image.Ic1 == CurveB_DTLicon ||
			Track[i].Image.Ic1 == CurveR_DTLicon ||
			Track[i].Image.Ic1 == Cul_DTLicon {
			if InBounds(r, c-1, numRows, numCols) {
				tr[Cell{r, c + 1}] = v
			}
			if InBounds(r-1, c, numRows, numCols) {
				bl[Cell{r - 1, c}] = v
			}
		}

	}

	//DDownRight
	//DDownleft
	for i := range Track {
		cell := Cell{Track[i].CurPosR, Track[i].CurPosC}
		_, inTL := tl[cell]
		_, inTR := tr[cell]
		_, inBL := bl[cell]
		_, inBR := br[cell]

		switch {
		case inTL && inBR:
			Track[i].Image.Ic3 = CornersTL_BRicon
		case inTR && inBL:
			Track[i].Image.Ic3 = CornersTR_BLicon
		case inTL:
			Track[i].Image.Ic3 = CornerTLicon
		case inBR:
			Track[i].Image.Ic3 = CornerBRicon
		case inBL:
			Track[i].Image.Ic3 = CornerBLicon
		case inTR:
			Track[i].Image.Ic3 = CornerTRicon
		}
	}
}

func InBounds(r, c, numRows, numCols int) bool {
	if r >= 0 && r < numRows && c >= 0 && c < numCols {
		return true
	}
	return false
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
