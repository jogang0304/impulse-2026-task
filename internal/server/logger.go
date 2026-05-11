package server

import (
	"log"

	"github.com/jogang0304/impulse-2026-task/internal/interfaces"
	"github.com/jogang0304/impulse-2026-task/internal/types"
)

type EventLogger struct{}

func MakeEventLoggerPtr() *EventLogger {
	return &EventLogger{}
}

func (l *EventLogger) LogBossFloor(e types.Event) {
	log.Printf("[%v] Player [%d] entered the boss's floor\n", e.Time, e.Player)
}

func (l *EventLogger) LogBossKill(e types.Event) {
	log.Printf("[%v] Player [%d] killed the boss\n", e.Time, e.Player)
}

func (l *EventLogger) LogDamage(e types.Event) {
	log.Printf("[%v] Player [%d] recieved [%d] of damage\n", e.Time, e.Player, e.ExtraParamInt)
}

func (l *EventLogger) LogDeath(e types.Event) {
	log.Printf("[%v] Player [%d] is dead\n", e.Time, e.Player)
}

func (l *EventLogger) LogEntrance(e types.Event) {
	log.Printf("[%v] Player [%d] entered the dungeon\n", e.Time, e.Player)
}

func (l *EventLogger) LogHeal(e types.Event) {
	log.Printf("[%v] Player [%d] has restored [%d] health\n", e.Time, e.Player, e.ExtraParamInt)
}

func (l *EventLogger) LogKill(e types.Event) {
	log.Printf("[%v] Player [%d] killed the monster\n", e.Time, e.Player)
}

func (l *EventLogger) LogLeave(e types.Event) {
	log.Printf("[%v] Player [%d] left the dungeon\n", e.Time, e.Player)
}

func (l *EventLogger) LogNextFloor(e types.Event) {
	log.Printf("[%v] Player [%d] went to the next floor\n", e.Time, e.Player)
}

func (l *EventLogger) LogPreviousFloor(e types.Event) {
	log.Printf("[%v] Player [%d] went to the previous floor\n", e.Time, e.Player)
}

func (l *EventLogger) LogProblem(e types.Event) {
	log.Printf("[%v] Player [%d] cannot continue due to [%s]\n", e.Time, e.Player, e.ExtraParamStr)
}

func (l *EventLogger) LogRegistration(e types.Event) {
	log.Printf("[%v] Player [%d] registered\n", e.Time, e.Player)
}

func (l *EventLogger) LogImpossibleMove(e types.Event) {
	log.Printf("[%v] Player [%d] makes impossible move [%d]\n", e.Time, e.Player, e.Type)
}

func (l *EventLogger) LogDisqualification(e types.Event) {
	log.Printf("[%v] Player [%d] is disqualified\n", e.Time, e.Player)
}

var _ interfaces.EventLogger = &EventLogger{}
