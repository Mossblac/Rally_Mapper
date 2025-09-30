package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// update this function to take a struct or map so the number of images are optional

func SetImageInCell(row, col int, imageName_Path []string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fyne.Do(func() {

			img := canvas.NewImageFromResource(StartUPIcon) //all three
			img.FillMode = canvas.ImageFillContain          //of these
			img.Translucency = 1.0                          //for each image

			cell := CellGrid[row][col]
			if cell == nil || len(cell.Objects) < 2 {
				return
			}
			// cell.Objects[0] = border, cell.Objects[1] = stack
			border := cell.Objects[0]
			stack, ok := cell.Objects[1].(*fyne.Container)
			if !ok {
				return
			}

			// stack.objects are your images, stack.object[0] is on bottom, the next on top would be stack.object[1] and so on
			if len(stack.Objects) > 0 {
				bg := stack.Objects[0]
				stack.Objects = []fyne.CanvasObject{bg}
			} else {
				stack.Objects = nil
			}

			border.Hide() // this hides the square with border before placing new image.

			stack.Add(img)  // you can add another image here with another stack.add(image) or to a specific layer by index: stack.object[index]
			stack.Refresh() // call refresh after adding each new image to stack

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
