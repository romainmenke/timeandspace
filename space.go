package timeandspace

import (
	"math"
)

// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Go

func (s TimeAndSpace) LatitudeLongitude() (float64, float64) {
	n := math.Pi - 2.0*math.Pi*float64(s.SY)/math.Exp2(float64(s.SZ))
	lat := 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	long := float64(s.SX)/math.Exp2(float64(s.SZ))*360.0 - 180.0
	return lat, long
}

func NewSpaces(lat float64, long float64) []TimeAndSpace {
	spaces := make([]TimeAndSpace, 0, 16)

	var z uint8
	for z = 0; z < uint8(cap(spaces)); z++ {
		spaces = append(spaces, newSpace(z, lat, long))
	}

	return spaces
}

func newSpace(z uint8, lat float64, long float64) TimeAndSpace {
	x := uint16(math.Floor((long + 180.0) / 360.0 * (math.Exp2(float64(z)))))
	y := uint16(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(z)))))
	return TimeAndSpace{
		SZ: z,
		SX: x,
		SY: y,
	}
}
