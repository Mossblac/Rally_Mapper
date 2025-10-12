package main

func CloneTrack(src []TrackCell) []TrackCell {
	if src == nil {
		return nil
	}
	dst := make([]TrackCell, len(src))
	copy(dst, src)

	for i := range dst {
		dst[i].Image = IconSet{
			Ic1: src[i].Image.Ic1,
			Ic2: src[i].Image.Ic2,
			Ic3: src[i].Image.Ic3,
			Ic4: src[i].Image.Ic4,
			Ic5: src[i].Image.Ic5,
		}
	}
	return dst
}
