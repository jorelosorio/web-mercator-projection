package geo

type BBox struct {
	Top, Left, Bottom, Right float64
}

// WithinXAxis constraint value into the axis X (Left, Right).
func (bb *BBox) WithinXAxis(value float64) float64 {
	return within(value, bb.Left, bb.Right)
}

// WithinYAxis constraint value into the axis Y (Top, Bottom)
func (bb *BBox) WithinYAxis(value float64) float64 {
	return within(value, bb.Top, bb.Bottom)
}

// within If the parameter value is out of the contained values, the edges of the bounding box will be returned,
// otherwhise value will be returned if is still inside.
func within(value float64, min float64, max float64) float64 {
	if value <= min {
		return min
	} else if value >= max {
		return max
	}

	return value
}
