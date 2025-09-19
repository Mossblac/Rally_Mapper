package commands

/*
	take count of each obstacle.

// take option selection of loop or linear

// linear has three columns (wide), there are three rows per obstacle to include. (start and end is always center cell)
// a linear course with 6 obstacles would be a grid of 3 x 18 - three wide, eighteen high.

//loop course has three cells square (9) to start, add 1 cell both x and y for each obstacle.
// a loop course with 6 obstacles would be 15 X 15

//different obstacles take up different amount of spaces. but max of two cells wide, and three cells tall.
*/

func GetNumRows_Cols(Ctype string, Ocount int) (numrows, numcols int) {
	if Ctype == "loop" {
		numrows = Ocount + 9
		numcols = Ocount + 9
	}
	if Ctype == "linear" {
		numrows = Ocount * 3
		numcols = 3
	}
	return numrows, numcols
}
