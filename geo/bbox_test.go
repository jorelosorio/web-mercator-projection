package geo

import (
	"fmt"
	"testing"
)

func TestValueOutsideOfXAxis(t *testing.T) {
	if WorldBBox.WithinXAxis(-190) != -180 {
		t.Error("Wrong edge, the returned value should be -180")
	}
}

func TestValueOutsideOfYAxis(t *testing.T) {
	stringValue := fmt.Sprintf("%.6f", WorldBBox.WithinYAxis(-90))
	if stringValue != "-85.051129" {
		t.Error("Wrong edge, the returned value should be -85.051129")
	}
}

func TestValueInsideOfXAxis(t *testing.T) {
	if WorldBBox.WithinXAxis(12.03) != 12.03 {
		t.Error("Wrong edge, the returned value should be 12.03")
	}
}

func TestValueInsideOfYAxis(t *testing.T) {
	if WorldBBox.WithinYAxis(-12.03) != -12.03 {
		t.Error("Wrong edge, the returned value should be -12.03")
	}
}
