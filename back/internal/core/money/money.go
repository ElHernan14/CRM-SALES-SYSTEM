package money

import "math"

func Round(value float64) float64 {
	return math.Round(value*100) / 100
}

func Add(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return Round(total)
}

func Subtract(left float64, right float64) float64 {
	return Round(left - right)
}

func Remaining(total float64, paid float64) float64 {
	remaining := Subtract(total, paid)
	if remaining < 0 {
		return 0
	}
	return remaining
}

func GreaterThan(left float64, right float64) bool {
	return Round(left) > Round(right)
}

func GreaterOrEqual(left float64, right float64) bool {
	return Round(left) >= Round(right)
}
