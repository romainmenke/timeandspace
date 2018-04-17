package timeandspace

import (
	"fmt"
	"time"
)

type TimeAndSpace struct {
	TZ uint8 `json:"t_z" gorm:"type:TINYINT; not null"`
	SZ uint8 `json:"s_z" gorm:"type:TINYINT; not null"`

	TX uint16 `json:"t_x" gorm:"type:SMALLINT; not null"`
	TY uint16 `json:"t_y" gorm:"type:SMALLINT; not null"`
	SX uint16 `json:"s_x" gorm:"type:SMALLINT; not null"`
	SY uint16 `json:"s_y" gorm:"type:SMALLINT; not null"`

	ForeignID string `json:"foreign_id" gorm:"type:varchar(255); not null"`
}

func (ts TimeAndSpace) String() string {
	return fmt.Sprintf("tz: %3d tx: %5d ty: %5d sz: %3d sx: %5d sy: %5d id: %s", ts.TZ, ts.TX, ts.TY, ts.SZ, ts.SX, ts.SY, ts.ForeignID)
}

func NewTimesAndSpaces(id string, t time.Time, lat float64, long float64) []TimeAndSpace {

	times := NewTimes(t)
	spaces := NewSpaces(lat, long)

	timesAndSpaces := make([]TimeAndSpace, 0, len(times)*len(spaces))

	for _, t := range times {
		for _, s := range spaces {

			timesAndSpaces = append(timesAndSpaces, TimeAndSpace{
				ForeignID: id,
				TZ:        t.TZ,
				TX:        t.TX,
				TY:        t.TY,
				SZ:        s.SZ,
				SX:        s.SX,
				SY:        s.SY,
			})

		}
	}

	return timesAndSpaces
}
