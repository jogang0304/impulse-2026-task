package types

type PlayerStatus string

const (
	StatusSuccess      = "SUCCESS"
	StatusFail         = "FAIL"
	StatusDisqualified = "DISQUAL"
)

type Player struct {
	Registered              bool
	CurrentFloor            int
	Lives                   int
	InDungeon               bool
	Status                  PlayerStatus
	TotalTime               CustomTime
	AverageMonsterFloorTime CustomTime
	BossFloorTime           CustomTime
}

func MakePlayer() *Player {
	return &Player{
		Registered:              false,
		CurrentFloor:            0,
		Lives:                   100,
		InDungeon:               false,
		Status:                  StatusFail,
		TotalTime:               MakeZeroTime(),
		AverageMonsterFloorTime: MakeZeroTime(),
		BossFloorTime:           MakeZeroTime(),
	}
}
