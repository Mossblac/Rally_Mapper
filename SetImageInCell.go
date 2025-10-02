package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

// update this function to take a struct or map so the number of images are optional- DisplayTrkImages, SetStart, PickNext

func SetImageInCell(row, col int, icons IconSet) {
	go func() {
		fyne.Do(func() {
			res := theme.CancelIcon() // "x" icon
			ex := canvas.NewImageFromResource(res)

			img := ResourceToIcon(icons.Ic1)
			secondimg := ResourceToIcon(icons.Ic2)
			thirdimg := ResourceToIcon(icons.Ic3)
			fourthimg := ResourceToIcon(icons.Ic4)

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

			if icons.Ic1 != nil {
				stack.Add(img)  // you can add another image here with another stack.add(image) or to a specific layer by index: stack.object[index]
				stack.Refresh() // call refresh after adding each new image to stack

				FadeInAnimate(img)
			} else {
				stack.Add(ex)
				stack.Refresh()
			}

			if icons.Ic2 != nil {
				stack.Add(secondimg)
				stack.Refresh()

				FadeInAnimate(secondimg)
			}

			if icons.Ic3 != nil {
				stack.Add(secondimg)
				stack.Refresh()

				FadeInAnimate(thirdimg)
			}

			if icons.Ic4 != nil {
				stack.Add(secondimg)
				stack.Refresh()

				FadeInAnimate(fourthimg)
			}

		})
	}()
}

func FadeInAnimate(img *canvas.Image) { //you can add the time as an input and have each image fade in one after another
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

func ResourceToIcon(static *fyne.StaticResource) *canvas.Image {
	img := canvas.NewImageFromResource(static)
	img.FillMode = canvas.ImageFillContain
	img.Translucency = 1.0

	return img
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
