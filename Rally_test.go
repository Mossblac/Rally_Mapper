package main

import "testing"

func TestUnmarshalTrack(t *testing.T) {
	Track, err := LoadTrack("./TSaves", "53464be4-9cee-437f-958e-c99c6e42d1df")
	if err != nil {
		t.Errorf("unable to Load Sample Track: %v", err)
	}

	if Track.ID != "53464be4-9cee-437f-958e-c99c6e42d1df" {
		t.Errorf("unmarhal error")
	}
}

func TestCompleteTrack(t *testing.T) {
	Track, err := LoadTrack("./TSaves", "53464be4-9cee-437f-958e-c99c6e42d1df")
	if err != nil {
		t.Errorf("unable to Load Sample Track: %v", err)
	}
	var HasStart bool
	var HasFinish bool
	Tdata := Track.TrackData
	for _, cell := range Tdata {
		if cell.Start {
			HasStart = true
		}
		if cell.Finish {
			HasFinish = true
		}
	}

	if !HasStart || !HasFinish {
		t.Errorf("No start or finish found")
	}

}
