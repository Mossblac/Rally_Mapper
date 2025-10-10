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
		if IsDTopRight(r, c) { //DUPRight
			if InBounds(r, c+1, numRows, numCols) && !IsDTopLeft(r, c+1) { //right of
				tl[Cell{r, c + 1}] = v
				if !OnTrack(r, c+1) {
					MakeTrack(r, c+1)
				}
			}
			if InBounds(r-1, c, numRows, numCols) && !IsDBottomRight(r-1, c) { // above
				br[Cell{r - 1, c}] = v
				if !OnTrack(r-1, c) {
					MakeTrack(r-1, c)
				}
			}
		}

		if IsDTopLeft(r, c) { //DUPleft
			if InBounds(r, c-1, numRows, numCols) && !IsDTopRight(r, c-1) { //left of
				tr[Cell{r, c - 1}] = v
				if !OnTrack(r, c-1) {
					MakeTrack(r, c-1)
				}
			}
			if InBounds(r-1, c, numRows, numCols) && !IsDBottomLeft(r-1, c) { //above
				bl[Cell{r - 1, c}] = v
				if !OnTrack(r-1, c) {
					MakeTrack(r-1, c)
				}
			}
		}

		if IsDBottomRight(r, c) { //DDownRight
			if InBounds(r+1, c, numRows, numCols) && !IsDTopRight(r+1, c) {
				tr[Cell{r + 1, c}] = v
				if !OnTrack(r+1, c) {
					MakeTrack(r+1, c)
				}
			}
			if InBounds(r, c+1, numRows, numCols) && !IsDBottomLeft(r, c+1) {
				bl[Cell{r, c + 1}] = v
				if !OnTrack(r, c+1) {
					MakeTrack(r, c+1)
				}
			}
		}

		if IsDBottomLeft(r, c) { //DDownleft
			if InBounds(r+1, c, numRows, numCols) && !IsDTopLeft(r+1, c) {
				tl[Cell{r + 1, c}] = v
				if !OnTrack(r+1, c) {
					MakeTrack(r+1, c)
				}
			}
			if InBounds(r, c-1, numRows, numCols) && !IsDBottomRight(r, c-1) {
				br[Cell{r, c - 1}] = v
				if !OnTrack(r, c-1) {
					MakeTrack(r, c-1)
				}
			}
		}

	}

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

func IsDTopRight(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic1
			switch icon {
			case StartAngleURicon:
				return true
			case StraightDTRicon:
				return true
			case StraightDBLicon:
				return true
			case CurveB_DTRicon:
				return true
			case CurveL_DTRicon:
				return true
			case CurveDTR_Bicon:
				return true
			case CurveDTR_Licon:
				return true
			case Cul_DTRicon:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func IsDTopLeft(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic1
			switch icon {
			case StartAngleULicon:
				return true
			case StraightDTLicon:
				return true
			case StraightDBRicon:
				return true
			case CurveDTL_Bicon:
				return true
			case CurveDTL_Ricon:
				return true
			case CurveB_DTLicon:
				return true
			case CurveR_DTLicon:
				return true
			case Cul_DTLicon:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func IsDBottomRight(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic1
			switch icon {
			case StraightDTLicon:
				return true
			case StraightDBRicon:
				return true
			case CurveDBR_Licon:
				return true
			case CurveDBR_Ticon:
				return true
			case CurveL_DBRicon:
				return true
			case CurveT_DBRicon:
				return true
			case Cul_DBRicon:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func IsDBottomLeft(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic1
			switch icon {
			case StraightDTRicon:
				return true
			case StraightDBLicon:
				return true
			case CurveDBL_Ricon:
				return true
			case CurveDBL_Ticon:
				return true
			case CurveR_DBLicon:
				return true
			case CurveT_DBLicon:
				return true
			case Cul_DBLicon:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func OnTrack(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			return true
		}
	}
	return false
}

func MakeTrack(r, c int) {
	Track[TrackLength+1] = TrackCell{
		CurPosR: r,
		CurPosC: c,
	}
}
