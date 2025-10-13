package main

//Obstacles - []string - list of selected obstacles.
//TrackLength- int - used to calculate obstacle and punch placement
//TrackT - bool - if true track is loop

//this will run with Track being fully build - obstacle icons use ic4, punches use ic5 - set to drop in.
// 0.4 seconds per square at full speed. - three squares to reach full speed from start, cul, curb/drop, or parking block
// - two squares to reach full speed from low-bridge

// at least 1 punch per obstacle.
// if cul has 2 or more rev after it, the cul gets a punch- more time reduced per rev.

/*
func ObstaclesAndPunches(){}
*/

/*
To evenly space checkpoints on a track of a given length — so that the distance between the start, each checkpoint, and the end is equal — you can follow this formula:

✅ Formula

If:

L = length of the track (in spaces)

N = number of checkpoints

Then the spacing between each point (including start and end) is:

spacing = L / (N + 1)


Then the position of each checkpoint is:

checkpoint_i = spacing × i     where i = 1 to N

💡 Why N + 1?

You're dividing the track into (N + 1) equal segments:

One segment from start to first checkpoint

N−1 segments between checkpoints

One segment from last checkpoint to end

So with 3 checkpoints, you're making 4 segments.

🧠 Example 1:

Track length = 120

Checkpoints = 3

Spacing:

spacing = 120 / (3 + 1) = 30


Checkpoints:

Checkpoint 1 = 30
Checkpoint 2 = 60
Checkpoint 3 = 90


Start = 0
End = 120

Segments: [0, 30, 60, 90, 120] — evenly spaced by 30.

🧠 Example 2:

Track length = 900

Checkpoints = 5

spacing = 900 / (5 + 1) = 150

Checkpoints: 150, 300, 450, 600, 750

⚠️ Note:

If L / (N + 1) doesn't give an integer, you'll get non-integer positions (e.g., 14.4). Depending on your system, you may need to:

Round to nearest if allowed

Use floating point positions

Disallow certain combinations
*/
