package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var ProgBarCalc *widget.ProgressBarInfinite

var Track []TrackCell

var GridBackground *canvas.Rectangle

var CurrentStop chan struct{}
var SpacingList []int
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
	Ic5 *fyne.StaticResource
}

type TrackCell struct {
	TrackInt int
	CurPosR  int
	CurPosC  int
	PrevMov  string
	Visited  int
	Image    IconSet
	Start    bool
	Finish   bool
	Cul      bool
	Rev      bool
	RevRef   int
	PTime    time.Duration
}
