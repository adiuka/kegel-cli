package workout

import "time"

type Plan struct {
	Reps		int
	Squeeze	time.Duration
	Rest		time.Duration
}

var Default = Plan{
	Reps:			10,
	Squeeze: 	10 * time.Second,
	Rest:			5 * time.Second,
}