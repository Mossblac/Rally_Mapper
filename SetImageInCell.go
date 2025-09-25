package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func SetImageInCell(CellGrid [][]*fyne.Container, row, col int, imageName_Path []string) {
	if path, ok := ImagePaths[imageName_Path[0]]; ok {
		path.FillMode = canvas.ImageFillContain
		CellGrid[row][col].Objects = []fyne.CanvasObject{path}
		CellGrid[row][col].Refresh()
	} else {
		image := canvas.NewImageFromFile(imageName_Path[1])
		ImagePaths[imageName_Path[0]] = image
		image.FillMode = canvas.ImageFillContain
		CellGrid[row][col].Objects = []fyne.CanvasObject{image}
		CellGrid[row][col].Refresh()
	}
}

func FadeInAnimate() {

}
