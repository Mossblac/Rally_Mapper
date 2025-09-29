package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// go
func SetImageInCell(row, col int, imageName_Path []string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fyne.Do(func() {
			// build image (replace StartUPIcon with what you need from imageName_Path)
			img := canvas.NewImageFromResource(StartUPIcon)
			img.FillMode = canvas.ImageFillContain
			img.Translucency = 1.0 // start invisible

			cell := CellGrid[row][col]
			if cell == nil || len(cell.Objects) < 2 {
				return
			}
			// cell.Objects[0] = border, cell.Objects[1] = stack
			stack, ok := cell.Objects[1].(*fyne.Container)
			if !ok {
				return
			}

			// optional: keep stack's background if present (index 0), clear others
			if len(stack.Objects) > 0 {
				bg := stack.Objects[0]
				stack.Objects = []fyne.CanvasObject{bg}
			} else {
				stack.Objects = nil
			}

			stack.Add(img)
			stack.Refresh()

			FadeInAnimate(img)
		})
	}()
}

func FadeInAnimate(img *canvas.Image) {
	fadeIn := canvas.NewColorRGBAAnimation(
		color.RGBA{A: 0}, color.RGBA{A: 255},
		500*time.Millisecond,
		func(c color.Color) {
			_, _, _, a := c.RGBA()
			img.Translucency = 1.0 - (float64(a) / 65535.0)
			img.Refresh()
		},
	)
	fadeIn.Start()
}

/*func AddImageAt(r, c int, img fyne.CanvasObject) {
	cell := CellGrid[r][c]
	// cell.Objects[1] is the stack
	stack, _ := cell.Objects[1].(*fyne.Container)
	stack.Add(img)
	stack.Refresh()
}*/

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
