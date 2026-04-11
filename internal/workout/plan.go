package workout

import "time"

type Plan struct {
	Name		string
	Reps		int
	Squeeze	time.Duration
	Rest		time.Duration
}

var Default = Plan{
	Name:			"Default",
	Reps:			10,
	Squeeze: 	10 * time.Second,
	Rest:			5 * time.Second,
}

var Levels = []Plan{
	{
		Name:			"Beginner",
		Reps:			5,
		Squeeze: 	5 * time.Second,
		Rest:			3 * time.Second,
	},
	{
		Name:			"Easy",
		Reps:			10,
		Squeeze: 	5 * time.Second,
		Rest:			3 * time.Second,
	},
	{
		Name:			"Standard",
		Reps:			10,
		Squeeze: 	10 * time.Second,
		Rest:			5 * time.Second,
	},
	{
		Name:			"Advanced",
		Reps:			15,
		Squeeze: 	10 * time.Second,
		Rest:			5 * time.Second,
	},
	{
		Name:			"Expert",
		Reps:			20,
		Squeeze: 	15 * time.Second,
		Rest:			5 * time.Second,
	},
}

func ForLevel(level int) Plan {
	if level < 0 {
		return Levels[0]
	}
	if level >= len(Levels) {
		return Levels[len(Levels)-1]
	}
	return Levels[level]
}