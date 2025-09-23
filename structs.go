package main

type Map struct {
	Cells []Track `json:"cells"`
}

type Track struct {
	Position  []int  `json:"position"`
	Previous  []int  `json:"previous"`
	TrackIcon string `json:"trackicon"`
	Items     Items  `json:"items"`
}

type Items struct {
	Punch    string `json:"punch"`
	Obstacle string `json:"obstacle"`
}
