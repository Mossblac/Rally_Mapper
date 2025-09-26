package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func SetImageInCell(CellGrid [][]*fyne.Container, row, col int, imageName_Path []string) {
	image := canvas.NewImageFromFile(imageName_Path[1])
	image.FillMode = canvas.ImageFillContain
	image.Translucency = 1.0
	CellGrid[row][col].Objects = []fyne.CanvasObject{image}
	CellGrid[row][col].Refresh() //this line needs to be within a fyne.DoAndWait() or fyne.Do() to work
	FadeInAnimate(image)
	image = nil

}

func FadeInAnimate(image *canvas.Image) {
	fadeIn := canvas.NewColorRGBAAnimation(
		color.RGBA{A: 0},
		color.RGBA{A: 255},
		time.Second*1,
		func(c color.Color) {
			_, _, _, a := c.RGBA()
			image.Translucency = 1.0 - (float64(a) / 65535.0)
			image.Refresh()
		},
	)

	fadeIn.Start()
}
