package interfaces

import "github.com/jogang0304/impulse-2026-task/internal/types"

type EventLogger interface {
	LogImpossibleMove(e types.Event)
	LogDisqualification(e types.Event)
	LogDeath(e types.Event)
	LogRegistration(e types.Event)
	LogEntrance(e types.Event)
	LogKill(e types.Event)
	LogNextFloor(e types.Event)
	LogPreviousFloor(e types.Event)
	LogBossFloor(e types.Event)
	LogBossKill(e types.Event)
	LogLeave(e types.Event)
	LogProblem(e types.Event)
	LogHeal(e types.Event)
	LogDamage(e types.Event)
}
