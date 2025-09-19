package main

import (
	"fyne.io/fyne/v2"
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
	x2              float32
	y2              float32
	cell_size       int
	cell_num_row    int
	cell_num_col    int
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

func (c Cell) Draw() error {
	// draw all four lines for each cell wall
	commands.Draw_line(c.x1, c.y1, c.x2, c.y2, c.window)
	return nil
}
