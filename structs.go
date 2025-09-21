package main

type Map struct {
	Punches struct {
		Amount   int    `json:"amount"`
		Position string `json:"position"`
		Facing   string `json:"facing"`
	} `json:"punches"`
}
