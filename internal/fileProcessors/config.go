package fileprocessors

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jogang0304/impulse-2026-task/internal/types"
)

func LoadConfig(filePath string) types.Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var c types.Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
