package database

import "github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"

var Durations []structs.Duration

func AddDuration() {
	Durations = append(Durations,
		structs.Duration{Seconds: 3600},
		structs.Duration{Seconds: 7200},
	)
}
