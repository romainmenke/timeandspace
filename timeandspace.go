package timeandspace

import (
	"time"
)

type TimeAndSpace struct {
	TimeX0 uint16 `json:"time_x_0" gorm:"type:SMALLINT; not null"`
	TimeY0 uint16 `json:"time_y_0" gorm:"type:SMALLINT; not null"`
	TimeY1 uint16 `json:"time_y_1" gorm:"type:SMALLINT; not null"`
	TimeY2 uint16 `json:"time_y_2" gorm:"type:SMALLINT; not null"`
	TimeY3 uint16 `json:"time_y_3" gorm:"type:SMALLINT; not null"`

	SpaceX0 uint16 `json:"space_x_0" gorm:"type:SMALLINT; not null"`
	SpaceY0 uint16 `json:"space_y_0" gorm:"type:SMALLINT; not null"`

	SpaceX1 uint16 `json:"space_x_1" gorm:"type:SMALLINT; not null"`
	SpaceY1 uint16 `json:"space_y_1" gorm:"type:SMALLINT; not null"`

	SpaceX2 uint16 `json:"space_x_2" gorm:"type:SMALLINT; not null"`
	SpaceY2 uint16 `json:"space_y_2" gorm:"type:SMALLINT; not null"`

	SpaceX3 uint16 `json:"space_x_3" gorm:"type:SMALLINT; not null"`
	SpaceY3 uint16 `json:"space_y_3" gorm:"type:SMALLINT; not null"`

	SpaceX4 uint16 `json:"space_x_4" gorm:"type:SMALLINT; not null"`
	SpaceY4 uint16 `json:"space_y_4" gorm:"type:SMALLINT; not null"`

	SpaceX5 uint16 `json:"space_x_5" gorm:"type:SMALLINT; not null"`
	SpaceY5 uint16 `json:"space_y_5" gorm:"type:SMALLINT; not null"`

	SpaceX6 uint16 `json:"space_x_6" gorm:"type:SMALLINT; not null"`
	SpaceY6 uint16 `json:"space_y_6" gorm:"type:SMALLINT; not null"`

	SpaceX7 uint16 `json:"space_x_7" gorm:"type:SMALLINT; not null"`
	SpaceY7 uint16 `json:"space_y_7" gorm:"type:SMALLINT; not null"`

	SpaceX8 uint16 `json:"space_x_8" gorm:"type:SMALLINT; not null"`
	SpaceY8 uint16 `json:"space_y_8" gorm:"type:SMALLINT; not null"`

	SpaceX9 uint16 `json:"space_x_9" gorm:"type:SMALLINT; not null"`
	SpaceY9 uint16 `json:"space_y_9" gorm:"type:SMALLINT; not null"`

	SpaceX10 uint16 `json:"space_x_10" gorm:"type:SMALLINT; not null"`
	SpaceY10 uint16 `json:"space_y_10" gorm:"type:SMALLINT; not null"`

	SpaceX11 uint16 `json:"space_x_11" gorm:"type:SMALLINT; not null"`
	SpaceY11 uint16 `json:"space_y_11" gorm:"type:SMALLINT; not null"`

	SpaceX12 uint16 `json:"space_x_12" gorm:"type:SMALLINT; not null"`
	SpaceY12 uint16 `json:"space_y_12" gorm:"type:SMALLINT; not null"`

	SpaceX13 uint16 `json:"space_x_13" gorm:"type:SMALLINT; not null"`
	SpaceY13 uint16 `json:"space_y_13" gorm:"type:SMALLINT; not null"`

	SpaceX14 uint16 `json:"space_x_14" gorm:"type:SMALLINT; not null"`
	SpaceY14 uint16 `json:"space_y_14" gorm:"type:SMALLINT; not null"`

	SpaceX15 uint16 `json:"space_x_15" gorm:"type:SMALLINT; not null"`
	SpaceY15 uint16 `json:"space_y_15" gorm:"type:SMALLINT; not null"`

	ForeignID string `json:"foreign_id" gorm:"type:varchar(255); not null"`
}

func NewTimeAndSpace(id string, t time.Time, lat float64, long float64) TimeAndSpace {

	time := NewTime(t)
	timeAndSpace := NewSpace(lat, long)

	timeAndSpace.TimeX0 = time.TimeX0
	timeAndSpace.TimeY0 = time.TimeY0
	timeAndSpace.TimeY1 = time.TimeY1
	timeAndSpace.TimeY2 = time.TimeY2
	timeAndSpace.TimeY3 = time.TimeY3

	return timeAndSpace
}
