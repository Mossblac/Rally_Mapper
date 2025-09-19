package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/Mossblac/Rally_Mapper/commands"
)

var Grid []Cell

type Cell struct {
	has_left_wall   bool
	has_right_wall  bool
	has_top_wall    bool
	has_bottom_wall bool
	x1              float32
	y1              float32
	cell_size       int
	window          fyne.Window
	visited         bool
}

func BuildGrid() error {
	// this will take the numrows and numcols from the input widget on the initial page.
	// first cell (top, left) is cell_num_row 0, cell_num_col 0
	// create a cell struct for each row x col and add to Grid slice
	// for each cell c.Draw()
	return nil
}

func (c Cell) Draw() (CT, CL, CB, CR *canvas.Line, err error) {
	// draw all four lines for each cell wall
	x2 := c.x1 + float32(c.cell_size)
	y2 := c.y1 + float32(c.cell_size)
	CT = commands.Draw_line(c.x1, c.y1, x2, c.y1, c.window) //top
	CL = commands.Draw_line(c.x1, c.y1, c.x1, y2, c.window) // left
	CB = commands.Draw_line(c.x1, y2, x2, y2, c.window)     // bottom
	CR = commands.Draw_line(x2, y2, x2, c.y1, c.window)     //right

	return CT, CL, CB, CR, nil
}
