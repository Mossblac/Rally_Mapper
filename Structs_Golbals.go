package main

import "fyne.io/fyne/v2"

var Track []TrackCell

var TrackLength int

var NumObstacles int
var RevCount int
var TrackT bool
var PreSetCuls []int

type IconSet struct {
	Ic1 *fyne.StaticResource
	Ic2 *fyne.StaticResource
	Ic3 *fyne.StaticResource
	Ic4 *fyne.StaticResource
}

type TrackCell struct {
	TrackInt int
	CurPosR  int
	CurPosC  int
	PrevMov  string
	Visited  bool
	Image    IconSet
	Start    bool
	Finish   bool
	Cul      bool
	Rev      bool
	RevRef   int
}
