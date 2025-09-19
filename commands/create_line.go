package commands

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func Draw_line(x1, y1, x2, y2 float32, win fyne.Window) (line *canvas.Line) {
	line = canvas.NewLine(color.White)
	line.StrokeWidth = 2
	line.Position1 = fyne.NewPos(x1, y1)
	line.Position2 = fyne.NewPos(x2, y2)
	return line
}
