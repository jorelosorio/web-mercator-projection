package geo

import (
	"fmt"
	"math"
	"testing"
)

func TestConvertFromDegreeToRadians(t *testing.T) {
	pi := DegToRad(180)
	if strPi := fmt.Sprintf("%.8f", pi); strPi != "3.14159265" {
		t.Error("Converting from 180 degrees to radians should be the value of PI")
	}
}

func TestConvertFromRadiansToDegree(t *testing.T) {
	pi := math.Floor(RadToDeg(3.1416))
	if pi != 180 {
		t.Error("Converting PI from radians to degrees should be the value of 180")
	}
}
