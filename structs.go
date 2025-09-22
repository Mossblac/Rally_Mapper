package main

type Map struct {
	cells []GridSquare `json:"cells"`
}

type GridSquare struct {
	Current      []int  `json:"current"`
	Previous     []int  `json:"prevoius"`
	Punch        string `json:"punch"`
	ObstacleType string `json:"obstacleType"`
}
