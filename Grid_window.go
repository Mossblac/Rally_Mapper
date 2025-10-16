// go
package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	CellGrid    [][]*fyne.Container
	zoomLevel   float32 = 1.0
	minZoom     float32 = 1.0
	maxZoom     float32 = 0.0
	gridLayout  *squareGridLayout
	gridObjects []fyne.CanvasObject
	scroll      *container.Scroll
)

type squareGridLayout struct {
	rows, cols int
	cellSize   float32
}

func (l *squareGridLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	if l.rows == 0 || l.cols == 0 || len(objs) == 0 {
		return
	}

	i := 0
	for r := 0; r < l.rows; r++ {
		for c := 0; c < l.cols; c++ {
			if i >= len(objs) {
				return
			}
			x := float32(c) * l.cellSize
			y := float32(r) * l.cellSize

			objs[i].Move(fyne.NewPos(x, y))
			objs[i].Resize(fyne.NewSize(l.cellSize, l.cellSize))

			if cont, ok := objs[i].(*fyne.Container); ok {
				if len(cont.Objects) >= 2 {
					border := cont.Objects[0]
					stack := cont.Objects[1]
					border.Move(fyne.NewPos(0, 0))
					border.Resize(fyne.NewSize(l.cellSize, l.cellSize))
					const m float32 = 1
					stack.Move(fyne.NewPos(m, m))
					stack.Resize(fyne.NewSize(l.cellSize-2*m, l.cellSize-2*m))
				}
			}
			i++
		}
	}
}

func (l *squareGridLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(l.cellSize*float32(l.cols), l.cellSize*float32(l.rows))
}

func calculateCellSize(winSize fyne.Size, rows, cols int) float32 {

	availableWidth := winSize.Width - 100
	availableHeight := winSize.Height - 100
	return float32(math.Min(
		float64(availableWidth)/float64(cols),
		float64(availableHeight)/float64(rows),
	))
}

func Grid_Widget(trackType string, numObstacles int, T TrackSave) {
	if CurrentStop != nil {
		close(CurrentStop)
	}
	CurrentStop = make(chan struct{})

	NumObstacles = numObstacles
	Track = nil
	TrackLength = 0
	PreSetCuls = nil
	SpacingList = nil
	PunchList = nil

	var numCols, numRows int
	if trackType == "loop" {
		TrackT = true
		numCols = 3 * numObstacles
		numRows = 3 * numObstacles
	} else {
		TrackT = false
		numCols = 3
		numRows = 6 * numObstacles
	}

	winSize := mainWin.Canvas().Size()
	baseCellSize := calculateCellSize(winSize, numRows, numCols)
	defaultZoomLevel := float32(1.0) // Default Zoom Level
	zoomLevel = defaultZoomLevel
	gridLayout = &squareGridLayout{
		rows:     numRows,
		cols:     numCols,
		cellSize: baseCellSize * defaultZoomLevel,
	}

	maxZoom = float32(math.Floor(math.Min(
		float64(winSize.Width-100),
		float64(winSize.Height-100),
	) / float64(baseCellSize)))

	CellGrid = make([][]*fyne.Container, numRows)
	gridObjects = make([]fyne.CanvasObject, 0, numRows*numCols)
	for r := 0; r < numRows; r++ {
		CellGrid[r] = make([]*fyne.Container, numCols)
		for c := 0; c < numCols; c++ {
			border := canvas.NewRectangle(color.Gray16{Y: 0x4000})
			content := canvas.NewRectangle(color.Black)
			content.SetMinSize(fyne.NewSize(40, 40))

			stack := container.NewStack()
			bg := canvas.NewRectangle(color.Black)
			stack.Add(bg)

			cellContainer := container.NewWithoutLayout(border, stack)
			CellGrid[r][c] = cellContainer
			gridObjects = append(gridObjects, cellContainer)
		}
	}
	GridBackground = canvas.NewRectangle(color.Black)
	gridWidth := gridLayout.cellSize * float32(gridLayout.cols)
	gridHeight := gridLayout.cellSize * float32(gridLayout.rows)
	GridBackground.Resize(fyne.NewSize(gridWidth, gridHeight))

	grid := container.New(gridLayout, gridObjects...)
	centeredGrid := container.NewCenter(container.NewWithoutLayout(GridBackground, grid))
	scroll = container.NewScroll(centeredGrid)
	bkg := canvas.NewImageFromFile("./images/backgroundcropped.jpg")
	bkg.FillMode = canvas.ImageFillStretch
	scrollStack := container.NewStack(bkg, scroll)

	zoomIn := widget.NewButton("+", func() {
		if zoomLevel < maxZoom {
			zoomLevel *= 1.25
			if zoomLevel > maxZoom {
				zoomLevel = maxZoom
			}
			gridLayout.cellSize = baseCellSize * zoomLevel

			GridBackground.Resize(fyne.NewSize(
				gridLayout.cellSize*float32(gridLayout.cols),
				gridLayout.cellSize*float32(gridLayout.rows),
			))
			grid.Refresh()
			scroll.Refresh()
		}
	})
	zoomOut := widget.NewButton("-", func() {
		if zoomLevel > minZoom {
			zoomLevel /= 1.25
			if zoomLevel < minZoom {
				zoomLevel = minZoom
			}
			gridLayout.cellSize = baseCellSize * zoomLevel

			GridBackground.Resize(fyne.NewSize(
				gridLayout.cellSize*float32(gridLayout.cols),
				gridLayout.cellSize*float32(gridLayout.rows),
			))
			grid.Refresh()
			scroll.Refresh()
		}
	})
	resetZoom := widget.NewButton("100%", func() {
		zoomLevel = defaultZoomLevel
		gridLayout.cellSize = baseCellSize * zoomLevel

		GridBackground.Resize(fyne.NewSize(
			gridLayout.cellSize*float32(gridLayout.cols),
			gridLayout.cellSize*float32(gridLayout.rows),
		))
		grid.Refresh()
		scroll.ScrollToTop()
		scroll.Refresh()
	})

	homeButton := widget.NewButton("Home", func() {
		saveWin.Hide()
		punchWin.Hide()
		SafeStop()
		HomeScreen()
	})
	runAgainButton := widget.NewButton("Re-Generate", func() {
		saveWin.Hide()
		punchWin.Hide()
		if !Loading {
			SafeStop()
			go Grid_Widget(trackType, numObstacles, none)
		} else {
			return
		}
	})

	PunchInfoButton := widget.NewButton("Punch Info", func() {
		saveWin.Hide()
		PunchInfo(T)
		punchWin.Show()
	})
	ClickedOnce = false
	SaveWindowButton := widget.NewButton("Save", func() {
		punchWin.Hide()
		if !Loading {
			if !ClickedOnce {
				SaveWindow(numObstacles)
				saveWin.Show()
				ClickedOnce = true
			} else {
				SavedWindow()
				saveWin.Show()
			}
		} else {
			return
		}
	})

	bottomButtons := container.NewHBox(
		zoomIn,
		zoomOut,
		resetZoom,
		layout.NewSpacer(),
		SaveWindowButton,
		layout.NewSpacer(),
		homeButton,
		layout.NewSpacer(),
		runAgainButton,
		layout.NewSpacer(),
		PunchInfoButton,
	)

	content := container.NewBorder(
		layout.NewSpacer(),
		bottomButtons,
		nil,
		nil,
		scrollStack,
	)

	mainWin.SetContent(content)

	for i := 0; i < numRows*numCols; i++ {
		cell := TrackCell{
			TrackInt: -1,
			CurPosR:  -1,
			CurPosC:  -1,
			PrevMov:  "",
			Visited:  0,
			Start:    false,
		}
		Track = append(Track, cell)
	}
	if Loading {
		PunchList = T.Punchlist
		DisplayTrkImages(CurrentStop, T.TrackData)
	} else {
		DeterminePath_setStart(CurrentStop, trackType, numRows, numCols)
	}
}

func SafeStop() {
	if CurrentStop != nil {
		close(CurrentStop)
		CurrentStop = nil
	}
}
