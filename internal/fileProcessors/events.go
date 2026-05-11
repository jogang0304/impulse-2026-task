package fileprocessors

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jogang0304/impulse-2026-task/internal/types"
)

func makeEventFromLine(line string) types.Event {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		log.Fatal("Invalid event format")
	}
	var e types.Event

	timeStr := strings.Trim(parts[0], "[]")
	parsedTime, err := time.Parse(time.TimeOnly, timeStr)
	if err != nil {
		log.Fatalf("Error parsing time: %v\n", err)
	}

	e.SetTime(parsedTime)

	e.Player, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error parsing player id: %v\n", err)
	}

	eventTypeId, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalf("Error parsing event type id: %v\n", err)
	}

	err = e.SetType(eventTypeId)
	if err != nil {
		log.Fatalf("Error setting event type id: %v\n", err)
	}

	var extraParam string
	if len(parts) > 3 {
		extraParam = parts[3]
		err = e.SetExtraParam(extraParam)
		if err != nil {
			log.Fatalf("Error setting extra parameter: %v \n", err)
		}
	}

	return e
}

func GetEventsStream(filePath string) <-chan types.Event {
	c := make(chan types.Event)

	go func() {
		defer close(c)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			e := makeEventFromLine(line)
			c <- e
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return c
}
