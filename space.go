package timeandspace

import (
	"math"
)

// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Go

func (s TimeAndSpace) LatitudeLongitude() (float64, float64) {
	n := math.Pi - 2.0*math.Pi*float64(s.SpaceY15)/math.Exp2(float64(15))
	lat := 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	long := float64(s.SpaceX15)/math.Exp2(float64(15))*360.0 - 180.0
	return lat, long
}

func NewSpace(lat float64, long float64) TimeAndSpace {
	s := TimeAndSpace{}

	s.SpaceX0, s.SpaceY0 = newSpaceForZ(0, lat, long)
	s.SpaceX1, s.SpaceY1 = newSpaceForZ(1, lat, long)
	s.SpaceX2, s.SpaceY2 = newSpaceForZ(2, lat, long)
	s.SpaceX3, s.SpaceY3 = newSpaceForZ(3, lat, long)
	s.SpaceX4, s.SpaceY4 = newSpaceForZ(4, lat, long)
	s.SpaceX5, s.SpaceY5 = newSpaceForZ(5, lat, long)
	s.SpaceX6, s.SpaceY6 = newSpaceForZ(6, lat, long)
	s.SpaceX7, s.SpaceY7 = newSpaceForZ(7, lat, long)
	s.SpaceX8, s.SpaceY8 = newSpaceForZ(8, lat, long)
	s.SpaceX9, s.SpaceY9 = newSpaceForZ(9, lat, long)
	s.SpaceX10, s.SpaceY10 = newSpaceForZ(10, lat, long)
	s.SpaceX11, s.SpaceY11 = newSpaceForZ(11, lat, long)
	s.SpaceX12, s.SpaceY12 = newSpaceForZ(12, lat, long)
	s.SpaceX13, s.SpaceY13 = newSpaceForZ(13, lat, long)
	s.SpaceX14, s.SpaceY14 = newSpaceForZ(14, lat, long)
	s.SpaceX15, s.SpaceY15 = newSpaceForZ(15, lat, long)

	return s
}

func newSpaceForZ(z uint8, lat float64, long float64) (uint16, uint16) {
	x := uint16(math.Floor((long + 180.0) / 360.0 * (math.Exp2(float64(z)))))
	y := uint16(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(z)))))
	return x, y
}
