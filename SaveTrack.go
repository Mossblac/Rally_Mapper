package main

func Save(TrackName string, numObstacles int) {
	save := TrackSave{
		Name:           TrackName,
		TimeToComplete: TotalTrackTime(),
		Ttype:          TrackT,
		TSize:          numObstacles,
		Punchlist:      PunchList,
		TrackData:      TrackClone,
	}
	Tsave := &save
	SaveNewTrack("./Saves", Tsave)
}
