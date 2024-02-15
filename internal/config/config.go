package config

import (
	"fmt"

	"github.com/rytsh/liz/file"
)

var FileAPI = file.New()

type Config struct {
	Input string `cfg:"input"`
	Sheet string `cfg:"sheet"`

	Map    map[string]map[string]string `cfg:"map"`
	Parse  map[string]Parse             `cfg:"parse"`
	Export []Export                     `cfg:"export"`
}

type Parse struct {
	Map  string   `cfg:"map"`
	Rows []string `cfg:"rows"`
}

type Export struct {
	Name     string `cfg:"name"`
	Template string `cfg:"template"`
	Output   string `cfg:"output"`
}

var AppConfig Config

func Load(fileName string) error {
	if err := FileAPI.Load(fileName, &AppConfig); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}
