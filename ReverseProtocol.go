package main

import "fyne.io/fyne/v2"

func ReverseProtocol(CellGrid [][]*fyne.Container, R, C, TrkI int) {
	Ccell := Trk[TrkI]
	_, ok := Ccell["Rev"].(bool)
	if !ok {
		RevCount += 1
		Ccell["Rev"] = true

		prev := Ccell["Previous"]
		switch prev {
		case "UP":
			Ccell["Previous"] = "Down"
		case "DUPLeft":
			Ccell["Previous"] = "DDownRight"
		case "DUPRight":
			Ccell["Previous"] = "DDownLeft"
		case "Right":
			Ccell["Previous"] = "Left"
		case "Left":
			Ccell["Previous"] = "Right"
		case "DDownRight":
			Ccell["Previous"] = "DUPLeft"
		case "DDownLeft":
			Ccell["Previous"] = "DUPRight"
		case "Down":
			Ccell["Previous"] = "UP"
		}

		PickNext(CellGrid, R, C, TrkI+1)
	}
	if ok {
		RevCount += 1
		//PrevCell := Trk[TrkI-1]
		//backtrack, ok := PrevCell["Position"].([]int)
		if ok {
			//BTR := backtrack[0]
			//BTC := backtrack[1]
		}
		//PickNext()
	}

}
