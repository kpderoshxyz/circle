package clock

import "fmt"

type Circle struct {
	diameter float32
}

func (circle *Circle) Radius() float32 {
	return circle.diameter / 2

}

type Point struct {
	x, y float32
}

// data structure to store the start x,y and end x,y of the line to be drawn
// Also hold the length of the line to be drawn. This will be calculated based on the circle passed in
type HourLine struct {
	End    Point
	Length float32
}

// Based on the hour figure out the angle at which the the hand needs to be drawn.
// We need multiple parameters to calculate the line
func GetHourLine(circle Circle, hour int) (*HourLine, error) {
	if hour > 24 || hour < 0 {
		return nil, fmt.Errorf("hour = %d argument is invalid", hour)
	}
	// Calculate endpoint point based on the hour integer
	return &HourLine{
		Length: (2 / 3) * (circle.Radius()),
	}, nil

}
