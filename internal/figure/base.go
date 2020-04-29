package figure

import (
	"math"
)

type (
	Square struct {
	}

	Triangle struct {
	}

	Rectangle struct {
	}

	Calculation interface {
		Perimeter(a ...float64) float64
		Area(a ...float64) float64
	}
)

func (s Square) Perimeter(a float64) float64 {
	var fourSide float64 = 4
	result := fourSide * a

	return result
}

func (s Square) Area(a float64) float64 {
	var secondPow float64 = 2
	result := math.Pow(secondPow, a)

	return result
}

func (s Triangle) Perimeter(a float64, b float64, c float64) float64 {
	result := a + b + c

	return result
}

func (s Triangle) Area(a float64, b float64, c float64) float64 {
	var second float64 = 2
	halfPer := (a + b + c) / second
	result := math.Sqrt(halfPer * (halfPer - a) * (halfPer - b) * (halfPer - c))

	return result
}

func (s Rectangle) Perimeter(r float64) float64 {
	var second float64 = 2
	result := second * math.Pi * r

	return result
}

func (s Rectangle) Area(r float32) float64 {
	var secondPow float64 = 2
	result := math.Pi * math.Pow(secondPow, float64(r))

	return result
}

