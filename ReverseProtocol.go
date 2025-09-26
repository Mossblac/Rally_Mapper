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
		TrkInt += 1
		PickNext(CellGrid, R, C, TrkI)
	}
	if ok {
		RevCount += 1
		TrkInt += 1
		PickNext(CellGrid, R, C, TrkI-1)
	}

}
