package types

import (
	"errors"
	"strconv"
	"time"
)

type EventType int

const (
	EventRegistration EventType = iota + 1
	EventEntrance
	EventKill
	EventNextFloor
	EventPreviousFloor
	EventBossFloor
	EventBossKill
	EventLeave
	EventProblem
	EventHeal
	EventDamage
	EventDisqual = iota + 31
	EventDeath
	EventImpossibleMove
)

type Event struct {
	Time            CustomTime
	Type            EventType
	Player          int
	ExtraParamStr   string
	ExtraParamInt   int
	ExtraParamIsInt bool
}

func (e *Event) SetTime(t time.Time) {
	e.Time.Time = t
}

func (e *Event) SetType(n int) error {
	switch n {
	case 1:
		e.Type = EventRegistration
	case 2:
		e.Type = EventEntrance
	case 3:
		e.Type = EventKill
	case 4:
		e.Type = EventNextFloor
	case 5:
		e.Type = EventPreviousFloor
	case 6:
		e.Type = EventBossFloor
	case 7:
		e.Type = EventBossKill
	case 8:
		e.Type = EventLeave
	case 9:
		e.Type = EventProblem
	case 10:
		e.Type = EventHeal
	case 11:
		e.Type = EventDamage
	case 31:
		e.Type = EventDisqual
	case 32:
		e.Type = EventDeath
	case 33:
		e.Type = EventImpossibleMove
	default:
		return errors.New("Unknown event id")
	}

	return nil
}

func (e *Event) SetExtraParam(s string) error {
	switch e.Type {
	case 0:
		return errors.New("event type should be set first")
	case 9:
		e.ExtraParamStr = s
		e.ExtraParamIsInt = false
	case 10, 11:
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		e.ExtraParamInt = n
		e.ExtraParamIsInt = true
	default:
		return errors.New("extra parameter is not needed")
	}

	return nil
}
