package main

import "fyne.io/fyne/v2"

var Track []TrackCell

var PossibleMoves []map[string][]int
var RevCount int
var TrackT bool

type IconSet struct {
	Ic1 *fyne.StaticResource
	Ic2 *fyne.StaticResource
	Ic3 *fyne.StaticResource
	Ic4 *fyne.StaticResource
}

type TrackCell struct {
	CurPosR int
	CurPosC int
	PrevMov string
	Visited bool
	Image   *fyne.StaticResource
	Start   bool
}
