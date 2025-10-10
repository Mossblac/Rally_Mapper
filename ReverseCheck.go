package main

func ReverseCheck() bool {
	for i := range TrackLength {
		if Track[i].Visited >= 4 {
			return true
		}
	}
	return false
}
