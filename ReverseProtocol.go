package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func ReverseProtocol(CellGrid [][]*fyne.Container, R, C, I int) {
	Ccell := Trk[I]
	_, ok := Ccell["Rev"].(bool)
	if !ok {
		fmt.Println("rev not found, adding rev")
		RevCount++
		Ccell["Rev"] = true

		prev := Ccell["Previous"]
		fmt.Printf("Original Previous value: %v", prev)
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
		fmt.Printf("Previous swapped to: %v", Ccell["Previous"])

		PickNext(CellGrid, R, C, I) // its all about the integer
	}
	if ok {
		fmt.Println("rev found, backtracking")
		RevCount++
		TrkInt++
		PickNext(CellGrid, R, C, I-1)
	}

}
