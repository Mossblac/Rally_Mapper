package main

type Config struct {
	Punches    int
	Curb_Nudge int
	Drop       int
	Hurdle     int
	Bonk       int
	Grab       int
	Catch      int
	Low_bridge int
	Tee_up     int
	Target     int
	D_Level    int
}

/* curb_nudges can be followed by a ramp
- ramps can lead to a drop
- a curb nudge can use a platform to end with a drop
- a hurdle is just a curb with no ramp or platform */
