package server

import (
	"log"

	"github.com/jogang0304/impulse-2026-task/internal/interfaces"
	"github.com/jogang0304/impulse-2026-task/internal/types"
)

type Server struct {
	Players map[int]*types.Player
	Config  types.Config
	Logger  interfaces.EventLogger
}

func MakeServer(c types.Config) *Server {
	var s Server
	s.Config = c
	s.Players = make(map[int]*types.Player)
	s.Logger = MakeEventLoggerPtr()

	return &s
}

func (s *Server) registerPlayer(p *types.Player) {
	p.Registered = true
}

func (s *Server) processOneEvent(e types.Event) {
	p, ok := s.Players[e.Player]
	if !ok {
		s.Players[e.Player] = types.MakePlayer()
		p = s.Players[e.Player]
	}

	if e.Type == types.EventRegistration {
		s.registerPlayer(p)
		s.Logger.LogRegistration(e)
		return
	}

	if !ok {
		s.Logger.LogDisqualification(e)
		return
	}

	if e.Type == types.EventEntrance {
		if p.InDungeon || p.Lives <= 0 {
			s.Logger.LogImpossibleMove(e)
		} else {
			p.InDungeon = true
			s.Logger.LogEntrance(e)
		}
		return
	}

	// All further events require player to be in dungeon and to be alive
	if !p.InDungeon && p.Lives > 0 {
		s.Logger.LogImpossibleMove(e)
		return
	} else if p.Lives <= 0 {
		s.Logger.LogDeath(e)
		return
	}

	switch e.Type {
	case types.EventLeave:
		p.InDungeon = false
		s.Logger.LogLeave(e)
	case types.EventBossFloor:
		s.Logger.LogBossFloor(e)
	case types.EventBossKill:
		s.Logger.LogBossKill(e)
	case types.EventDamage:
		s.Logger.LogDamage(e)
		p.Lives = max(0, p.Lives-e.ExtraParamInt)
	case types.EventHeal:
		s.Logger.LogHeal(e)
		p.Lives = min(100, p.Lives+e.ExtraParamInt)
	case types.EventKill:
		s.Logger.LogKill(e)
	case types.EventNextFloor:
		s.Logger.LogNextFloor(e)
		p.CurrentFloor++
	case types.EventPreviousFloor:
		s.Logger.LogPreviousFloor(e)
		p.CurrentFloor--
	case types.EventProblem:
		s.Logger.LogProblem(e)
	default:
		log.Fatalf("Unknown event type: %d\n", e.Type)
	}
}

func (s *Server) ProcessEventsFromChan(ch <-chan types.Event) {
	for e := range ch {
		s.processOneEvent(e)
	}
}
