package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/Mossblac/Rally_Mapper/commands"
)

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

// test inputs, will come from homemenu select. (if no input, display error)
//numrows is 18(tall), numcols is 3(wide)

func CreateCellsList() [][]Cell {
	numrows, numcols := commands.GetNumRows_Cols("linear", 6)

	Grid := make([][]Cell, numrows)
	for i := range Grid {
		Grid[i] = make([]Cell, numcols)
	}
	return Grid
}

func BuildGrid() {

	test_cell := Cell{
		has_left_wall:   true,
		has_right_wall:  true,
		has_top_wall:    true,
		has_bottom_wall: true,
		x1:              1,
		y1:              1,
		cell_size:       20,
		window:          course,
		visited:         false,
	}

	CT, CL, CB, CR, err := test_cell.Draw() //draw each cell in Grid cell list
	if err != nil {
		fmt.Printf("error creating lines %v", err)
	}

	courseContainer := container.NewWithoutLayout(CT, CL, CB, CR)
	course.SetContent(courseContainer)
	course.Resize(fyne.NewSize(float32(test_cell.cell_size*10), float32(test_cell.cell_size*10)))

	course.Show()
}

func (c Cell) Draw() (CT, CL, CB, CR *canvas.Line, err error) {
	// draw all four lines for each cell wall
	x2 := c.x1 + float32(c.cell_size)
	y2 := c.y1 + float32(c.cell_size)
	if c.has_top_wall {
		CT = commands.Draw_line(c.x1, c.y1, x2, c.y1, c.window) //top
	}
	if c.has_left_wall {
		CL = commands.Draw_line(c.x1, c.y1, c.x1, y2, c.window) // left
	}
	if c.has_bottom_wall {
		CB = commands.Draw_line(c.x1, y2, x2, y2, c.window) // bottom
	}
	if c.has_right_wall {
		CR = commands.Draw_line(x2, y2, x2, c.y1, c.window) //right
	}

	return CT, CL, CB, CR, nil
}
