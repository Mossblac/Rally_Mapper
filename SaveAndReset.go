package main

func ResetAndTryAgain(numRows, numCols int) {
	Track = Track[:0]

	for i := 0; i < numRows*numCols; i++ {
		cell := TrackCell{
			CurPosR: -1,
			CurPosC: -1,
			PrevMov: "",
			Visited: false,
			Start:   false,
		}
		Track = append(Track, cell)
	}

	if TrackT {
		DeterminePath_setStart("loop", numRows, numCols)
	} else {
		DeterminePath_setStart("linear", numRows, numCols)
	}
}
