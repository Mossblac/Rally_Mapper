package main

import (
	"fyne.io/fyne/v2"
)

func ResetAndTryAgain() {
	RevCount = 0
	if TrackT {
		fyne.Do(func() {
			Grid_Widget("loop", NumObstacles)
		})
	} else {
		fyne.Do(func() {
			Grid_Widget("linear", NumObstacles)
		})
	}
}

/*
// go
type Track struct {
    ID             string    `json:"id"`
    Name           string    `json:"name"`
    CreatedAt      time.Time `json:"createdAt"`
    TimeToComplete time.Duration `json:"timeToComplete"`
    Rows           int       `json:"rows"`
    Cols           int       `json:"cols"`
    ObstaclesCount int       `json:"obstaclesCount"`
    PunchesCount   int       `json:"punchesCount"`
    // ... your nested structs, SVG refs, etc.
}

type CatalogEntry struct {
    ID             string        `json:"id"`
    Name           string        `json:"name"`
    CreatedAt      time.Time     `json:"createdAt"`
    TimeToComplete time.Duration `json:"timeToComplete"`
    Rows           int           `json:"rows"`
    Cols           int           `json:"cols"`
    ObstaclesCount int           `json:"obstaclesCount"`
    PunchesCount   int           `json:"punchesCount"`
}

type Catalog struct {
    Entries []CatalogEntry `json:"entries"`
}

func (c *Catalog) Add(e CatalogEntry) {
    c.Entries = append(c.Entries, e) // immutable: no updates
}

// go
func atomicWriteJSON(path string, v any, perm fs.FileMode) error {
    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0o755); err != nil { return err }
    tmp, err := os.CreateTemp(dir, ".tmp-*")
    if err != nil { return err }
    enc := json.NewEncoder(tmp)
    enc.SetIndent("", "  ")
    if err := enc.Encode(v); err != nil { tmp.Close(); os.Remove(tmp.Name()); return err }
    if err := tmp.Sync(); err != nil { tmp.Close(); os.Remove(tmp.Name()); return err }
    if err := tmp.Close(); err != nil { os.Remove(tmp.Name()); return err }
    if err := os.Chmod(tmp.Name(), perm); err != nil { os.Remove(tmp.Name()); return err }
    return os.Rename(tmp.Name(), path)
}


// go
func SaveNewTrack(baseDir string, t *Track) error {
    now := time.Now().UTC()
    t.ID = uuid.NewString()
    t.CreatedAt = now

    trackPath := filepath.Join(baseDir, "tracks", t.ID+".json")
    if err := atomicWriteJSON(trackPath, t, 0o644); err != nil {
        return err
    }

    catalogPath := filepath.Join(baseDir, "catalog.json")
    var cat Catalog
    if b, err := os.ReadFile(catalogPath); err == nil {
        _ = json.Unmarshal(b, &cat)
    }
    cat.Add(CatalogEntry{
        ID:             t.ID,
        Name:           t.Name,
        CreatedAt:      t.CreatedAt,
        TimeToComplete: t.TimeToComplete,
        Rows:           t.Rows,
        Cols:           t.Cols,
        ObstaclesCount: t.ObstaclesCount,
        PunchesCount:   t.PunchesCount,
    })
    return atomicWriteJSON(catalogPath, &cat, 0o644)
}



*/

/*
// go
type Catalog struct {
    Entries []CatalogEntry `json:"entries"`
}

type CatalogEntry struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Rows int    `json:"rows"`
    Cols int    `json:"cols"`
    // ... other fields you want to display
}

func loadCatalog(path string) (Catalog, error) {
    var cat Catalog
    b, err := os.ReadFile(path)
    if err != nil {
        return cat, err
    }
    err = json.Unmarshal(b, &cat)
    return cat, err
}

func loadTrack(baseDir, id string) (*Track, error) {
    p := filepath.Join(baseDir, "tracks", id+".json")
    b, err := os.ReadFile(p)
    if err != nil {
        return nil, err
    }
    var t Track
    if err := json.Unmarshal(b, &t); err != nil {
        return nil, err
    }
    return &t, nil
}

func makeCatalogList(cat Catalog, onSelect func(CatalogEntry)) fyne.CanvasObject {
    data := make([]string, len(cat.Entries))
    for i, e := range cat.Entries {
        data[i] = fmt.Sprintf("%s (%dx%d)", e.Name, e.Rows, e.Cols)
    }

    list := widget.NewList(
        func() int { return len(data) },
        func() fyne.CanvasObject { return widget.NewLabel("item") },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(data[i])
        },
    )

    list.OnSelected = func(i widget.ListItemID) {
        if i >= 0 && i < len(cat.Entries) {
            onSelect(cat.Entries[i])
        }
    }
    return list
}


// go
func makeMainUI(baseDir string, w fyne.Window) (fyne.CanvasObject, error) {
    cat, err := loadCatalog(filepath.Join(baseDir, "catalog.json"))
    if err != nil {
        return nil, err
    }

    preview := widget.NewLabel("Select a track to preview")
    list := makeCatalogList(cat, func(e CatalogEntry) {
        t, err := loadTrack(baseDir, e.ID)
        if err != nil {
            dialog.ShowError(err, w)
            return
        }
        // Replace with your real drawing/rendering:
        preview.SetText(fmt.Sprintf("Track: %s\nSize: %dx%d", t.Name, t.Rows, t.Cols))
    })

    split := container.NewHSplit(list, container.NewVScroll(preview))
    split.SetOffset(0.35)
    return split, nil
}


*/
