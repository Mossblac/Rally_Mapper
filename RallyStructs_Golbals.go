package main

import "fyne.io/fyne/v2"

//var Trk []TrackCell

var Trk []map[string]interface{}
var PossibleMoves []map[string][]int
var RevCount int
var TrkInt int

type IconSet struct {
	Ic1 *fyne.StaticResource
	Ic2 *fyne.StaticResource
	Ic3 *fyne.StaticResource
	Ic4 *fyne.StaticResource
}

type TrackCell struct {
	CurPosX int
	CurPosY int
	PrevMov string
	Visited bool
}
