package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var ProgBarCalc *widget.ProgressBarInfinite

var Track []TrackCell
var TrackName string

var GridBackground *canvas.Rectangle

var CurrentStop chan struct{}
var SpacingList []int
var TrackLength int

var NumObstacles int
var RevCount int
var TrackT bool
var PreSetCuls []int
var PunchList []int

var ClickedOnce bool
var Loading bool
var none TrackSave
var LoadedTrk TrackSave

type IconSet struct {
	Ic1 *fyne.StaticResource `json:"ic1"`
	Ic2 *fyne.StaticResource `json:"ic2"`
	Ic3 *fyne.StaticResource `json:"ic3"`
	Ic4 *fyne.StaticResource `json:"ic4"`
	Ic5 *fyne.StaticResource `json:"ic5"`
}

type TrackCell struct {
	TrackInt int           `json:"trackInt"`
	CurPosR  int           `json:"curPosR"`
	CurPosC  int           `json:"curPosC"`
	PrevMov  string        `json:"prevMov"`
	Visited  int           `json:"visited"`
	Image    IconSet       `json:"image"`
	Start    bool          `json:"start"`
	Finish   bool          `json:"finish"`
	Cul      bool          `json:"cul"`
	Rev      bool          `json:"rev"`
	RevRef   int           `json:"revRef"`
	PTime    time.Duration `json:"pTime"`
}
