package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var crossRes = theme.CancelIcon()

func SetImageInCell(row, col int, imageName_Path []string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fyne.Do(func() {
			image := canvas.NewImageFromResource(StraightUPicon)
			image.FillMode = canvas.ImageFillContain
			image.Translucency = 1.0 // Start invisible

			CellGrid[row][col].Objects = []fyne.CanvasObject{image}
			CellGrid[row][col].Refresh()

			FadeInAnimate(image)
		})
	}()
}

// Call it like:
//SetImageInCell(CellGrid, row, col, imageName_Path, time.Duration(index)*200*time.Millisecond)

func FadeInAnimate(image *canvas.Image) {
	fadeIn := canvas.NewColorRGBAAnimation(
		color.RGBA{A: 0},
		color.RGBA{A: 255},
		time.Millisecond*500,
		func(c color.Color) {
			_, _, _, a := c.RGBA()
			image.Translucency = 1.0 - (float64(a) / 65535.0)
			image.Refresh()
		},
	)

	fadeIn.Start()
}

/*func DropInAnimate(image *canvas.Image, container *fyne.Container) {
	// Start above the container
	originalPos := image.Position()
	image.Move(fyne.NewPos(originalPos.X, originalPos.Y-100))
	image.Translucency = 0.0 // Start visible

	// Animate dropping down
	dropAnimation := canvas.NewPositionAnimation(
		fyne.NewPos(originalPos.X, originalPos.Y-100),
		originalPos,
		time.Millisecond*300,
		func(pos fyne.Position) {
			image.Move(pos)
			image.Refresh()
		},
	)

	dropAnimation.Start()
}*/
