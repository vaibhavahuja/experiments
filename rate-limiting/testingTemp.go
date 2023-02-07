package main

import (
	"fmt"
	"time"
)

type CustomTime struct {
	Year   int
	Month  string
	Day    int
	Hour   int
	Minute int
	Second int
}

//todo sure I can do this? but I am sure there exists an easier way?

func GetCustomTimeUptoUnit(time time.Time, unit string) CustomTime {
	switch unit {
	case "MINUTE":
		return CustomTime{
			Year:   time.Year(),
			Month:  time.Month().String(),
			Day:    time.Day(),
			Hour: time.Hour(),
			Minute: time.Minute(),
		}
	case "DAY":
		return CustomTime{
			Year:  time.Year(),
			Month: time.Month().String(),
			Day:   time.Day(),
		}
	}
	return CustomTime{}
}

func main() {
	currentTime := time.Now()
	fmt.Println(GetCustomTimeUptoUnit(currentTime, "MINUTE"))
	fmt.Println(GetCustomTimeUptoUnit(currentTime.Add(5*time.Minute), "MINUTE"))

}
