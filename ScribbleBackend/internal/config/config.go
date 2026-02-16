package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ServerPort uint16
}

// Load reads SERVER_PORT from the environment and returns a Config.
// It returns an error if SERVER_PORT is missing or cannot be parsed as a uint16.
func Load() (Config, error) {
	s := os.Getenv("SERVER_PORT")
	if s == "" {
		return Config{}, fmt.Errorf("SERVER_PORT not set")
	}
	v, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return Config{}, fmt.Errorf("invalid SERVER_PORT %q: %w", s, err)
	}
	return Config{ServerPort: uint16(v)}, nil
}
