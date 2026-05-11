package types

import (
	"strings"
	"time"
)

type CustomTime struct {
	Time time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct *CustomTime) String() string {
	return ct.Time.Format(time.TimeOnly)
}

func MakeZeroTime() CustomTime {
	return CustomTime{
		Time: time.Unix(0, 0),
	}
}

type Config struct {
	Floors   int        `json:"Floors"`
	Monsters int        `json:"Monsters"`
	OpenAt   CustomTime `json:"OpenAt"`
	Duration int        `json:"Duration"`
}
