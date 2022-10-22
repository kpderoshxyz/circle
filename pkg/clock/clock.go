package clock

import (
	"fmt"
	"math"

	"github.com/kpderoshxyz/circle/pkg/angle"
)

const DEGREES_IN_CIRCLE = 360
const DEGREES_PER_HOUR = (DEGREES_IN_CIRCLE / 12) // How many degrees per hour
const NOON_DEGREES = DEGREES_IN_CIRCLE - 90       // Degrees of noon/12

type Circle struct {
	diameter float64
}

func (circle *Circle) Radius() float64 {
	return circle.diameter / 2

}

type Point struct {
	x, y float64
}

// data structure to store the start x,y and end x,y of the line to be drawn
// Also hold the length of the line to be drawn. This will be calculated based on the circle passed in
type HourLine struct {
	End    *Point
	Length float64
}

type Time struct {
	Hour   float64
	Minute float64
}

// Based on the hour figure out the angle at which the the hand needs to be drawn.
// We need multiple parameters to calculate the line
func GetHourLine(circle Circle, time Time) (*HourLine, error) {
	if time.Hour >= 24 || time.Hour < 0 {
		return nil, fmt.Errorf("hour = %x argument is invalid, must be integer ---> 0 < hour <= 24", time.Hour)
	}
	// Calculate the angle based on the hour value
	// Calculate the number of degrees based on the hour.
	var initialDegrees float64 = 270 + (DEGREES_PER_HOUR * time.Hour)
	if initialDegrees >= 360 {
		initialDegrees = initialDegrees - 360
	}
	// Add in ratio of how many minutes out of 60 there is.
	var degrees float64 = initialDegrees + (DEGREES_PER_HOUR)*(time.Minute/60)
	var radians float64 = angle.DegToRadians(degrees)
	var xCoord float64 = math.Cos(radians) * circle.Radius() * (1 / 2)
	var yCoord float64 = math.Sin(radians) * circle.Radius() * (1 / 2)

	// Calculate endpoint point based on the hour integer
	return &HourLine{
		Length: (1 / 2) * (circle.Radius()),
		End: &Point{
			x: xCoord,
			y: yCoord,
		},
	}, nil
}
