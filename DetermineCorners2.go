package main

func IsDTopRight2(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic2
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

func IsDTopLeft2(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic2
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

func IsDBottomRight2(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic2
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

func IsDBottomLeft2(r, c int) bool {
	for i := range Track {
		if r == Track[i].CurPosR && c == Track[i].CurPosC {
			icon := Track[i].Image.Ic2
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
