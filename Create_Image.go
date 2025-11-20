package main

import (
	"fmt"
	"image/png"
	"os"
	"time"
)

func CreateImage() {

	//Grid.Refresh()

	time.Sleep(50 * time.Millisecond)

	Timg := mainWin.Canvas().Capture()

	file, err := os.Create("test_img.png")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}

	defer file.Close()

	err = png.Encode(file, Timg)
	if err != nil {
		fmt.Printf("error encoding image: %v", err)
	} else {
		fmt.Println("image saved")
	}

}
