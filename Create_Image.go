package main

import (
	"fmt"
	"image/png"
	"os"
)

func CreateImage(trackname string) {

	mainWin.Canvas().Refresh(mainWin.Content())

	img := mainWin.Canvas().Capture()
	if img == nil {
		fmt.Println("Capture returned nil")
		return
	}

	f, err := os.Create(trackname + ".png")
	if err != nil {
		fmt.Println("create error:", err)
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Println("encode error:", err)
	} else {
		fmt.Println("image saved")
	}
}
