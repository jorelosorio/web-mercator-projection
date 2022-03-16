package geo

// Keep the `value` within `min` and `max` edges
func withinMinMax(value float64, min float64, max float64) float64 {
	if value <= min {
		return min
	} else if value >= max {
		return max
	}

	return value
}
