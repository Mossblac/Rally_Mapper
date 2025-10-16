package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func SetTrackImageInCell(row, col int, icons IconSet) {
	go func() {
		fyne.Do(func() {
			res := theme.CancelIcon() // "x" icon
			ex := canvas.NewImageFromResource(res)

			img := ResourceToIcon(icons.Ic1)
			secondimg := ResourceToIcon(icons.Ic2)
			thirdimg := ResourceToIcon(icons.Ic3)
			fourthimg := ResourceToIcon(icons.Ic4)
			fifthimg := ResourceToIcon(icons.Ic5)

			cell := CellGrid[row][col]
			if cell == nil || len(cell.Objects) < 2 {
				return
			}

			border := cell.Objects[0]
			stack, ok := cell.Objects[1].(*fyne.Container)
			if !ok {
				return
			}

			if len(stack.Objects) > 0 {
				bg := stack.Objects[0]
				stack.Objects = []fyne.CanvasObject{bg}
			} else {
				stack.Objects = nil
			}

			border.Hide()

			if icons.Ic1 != nil {
				stack.Add(img)
				stack.Refresh()

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
				stack.Add(thirdimg)
				stack.Refresh()

				FadeInAnimate(thirdimg)
			}

			if icons.Ic4 != nil {
				stack.Add(fourthimg)
				stack.Refresh()

				DropInAnimate(fourthimg, time.Millisecond*400)
			}

			if icons.Ic5 != nil {
				stack.Add(fifthimg)
				stack.Refresh()

				DropInAnimate(fifthimg, time.Millisecond*700)
			}

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

func ResourceToIcon(static *fyne.StaticResource) *canvas.Image {
	img := canvas.NewImageFromResource(static)
	img.FillMode = canvas.ImageFillContain
	img.Translucency = 1.0

	return img
}

func DropInAnimate(image *canvas.Image, delay time.Duration) {
	originalPos := image.Position()
	image.Move(fyne.NewPos(originalPos.X, originalPos.Y-100))
	image.Translucency = 0.0

	dropAnimation := canvas.NewPositionAnimation(
		fyne.NewPos(originalPos.X, originalPos.Y-100),
		originalPos,
		delay,
		func(pos fyne.Position) {
			image.Move(pos)
			image.Refresh()
		},
	)

	dropAnimation.Start()
}
