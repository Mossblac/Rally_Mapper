// go
package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

	Grid = container.New(gridLayout, gridObjects...)
	centeredGrid := container.NewCenter(container.NewWithoutLayout(GridBackground, Grid))
	scroll = container.NewScroll(centeredGrid)
	res := fyne.NewStaticResource("images/menuBkG.jpg", resourceMenuBkGJpgData)
	bkg := canvas.NewImageFromResource(res)
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
			Grid.Refresh()
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
			Grid.Refresh()
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
		Grid.Refresh()
		scroll.ScrollToTop()
		scroll.Refresh()
	})

	homeButton := widget.NewButton("Home", func() {
		ClickedOnce = false
		punchWin.Hide()
		SafeStop()
		HomeScreen()
	})
	runAgainButton := widget.NewButton("Re-Generate", func() {
		punchWin.Hide()
		if !Loading {
			SafeStop()
			go Grid_Widget(trackType, numObstacles, none)
		} else {
			return
		}
	})

	PunchInfoButton := widget.NewButton("Punch Info", func() {
		PunchInfo(T)
		punchWin.Show()
	})

	maxEntryLength := 25
	SaveName := widget.NewEntry()
	SaveName.PlaceHolder = "Enter Track Name"
	SaveName.MultiLine = false
	SaveName.Wrapping = fyne.TextWrapOff
	SaveName.OnChanged = func(name string) {
		if len(name) > maxEntryLength {
			SaveName.SetText(name[:maxEntryLength])
		}
		TrackName = name
	}

	var SavesToAppear *fyne.Container
	var AlreadySavedMessage *fyne.Container

	SaveButton := widget.NewButton("                 Save                 ", func() {
		Save(TrackName, numObstacles)
		CreateImage()
		SavesToAppear.Hide()

	})

	SaveBkg := canvas.NewRectangle(color.RGBA{R: 54, G: 1, B: 63, A: 255})
	SaveBkg.CornerRadius = 20
	SaveBWithEntry := container.NewVBox(SaveName, SaveButton)
	SaveBkgWithBut := container.NewStack(SaveBkg, SaveBWithEntry)
	CentSBwithEntry := container.NewCenter(SaveBkgWithBut)

	SavesToAppear = container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		CentSBwithEntry,
	)

	SavesToAppear.Hide()

	AlSavedClose := widget.NewButton("     close     ", func() {
		AlreadySavedMessage.Hide()
	})

	AlSavedText := widget.NewLabel("     Track Saved     ")

	SavedBkg := canvas.NewRectangle(color.RGBA{R: 54, G: 1, B: 63, A: 255})
	SavedBkg.CornerRadius = 20
	AlSavedButAndText := container.NewVBox(AlSavedText, AlSavedClose)
	AlsavedBkgandBut := container.NewStack(SavedBkg, AlSavedButAndText)
	CentAlSavedButandText := container.NewCenter(AlsavedBkgandBut)

	AlreadySavedMessage = container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		CentAlSavedButandText,
	)

	AlreadySavedMessage.Hide()

	ClickedOnce = false
	SaveWindowButton := widget.NewButton("Save", func() {
		punchWin.Hide()
		if !Loading {
			if !ClickedOnce {
				SavesToAppear.Show()
				ClickedOnce = true
			} else {
				AlreadySavedMessage.Show()
			}
		} else {
			return
		}
	})

	work := theme.SettingsIcon()
	Working := canvas.NewImageFromResource(work)
	Working.FillMode = canvas.ImageFillContain
	Working.SetMinSize(fyne.NewSize(45, 45))
	Working.Show()

	iconBackground := canvas.NewRectangle(color.Black)
	iconBackground.SetMinSize(fyne.NewSize(50, 50))
	iconBackground.CornerRadius = 20

	workingStack := container.NewStack(iconBackground, Working)

	workingBox := container.NewVBox(layout.NewSpacer(), workingStack)
	canvas.Refresh(Working)

	bottomButtons := container.NewHBox(
		zoomIn,
		zoomOut,
		resetZoom,
		workingBox,
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

	winstack := container.NewStack(content, SavesToAppear, AlreadySavedMessage)

	mainWin.SetContent(winstack)

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
		LoadedTrk = T
		PunchList = T.Punchlist
		DisplayTrkImages(CurrentStop, T.TrackData, Working)
	} else {
		DeterminePath_setStart(CurrentStop, trackType, numRows, numCols, Working)
	}
}

func SafeStop() {
	if CurrentStop != nil {
		close(CurrentStop)
		CurrentStop = nil
	}
}
