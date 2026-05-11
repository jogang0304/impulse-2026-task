package main

import (
	fileprocessors "github.com/jogang0304/impulse-2026-task/internal/fileProcessors"
	"github.com/jogang0304/impulse-2026-task/internal/server"
)

func main() {
	c := fileprocessors.LoadConfig("config.json")
	ch := fileprocessors.GetEventsStream("events")

	s := server.MakeServer(c)

	s.ProcessEventsFromChan(ch)
}
