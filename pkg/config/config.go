package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Conf stores Hype configuration.
type Conf struct {
	// Editor specifies the default editor.
	Editor string
	// Port specifies the port on which Hype listens.
	Port string
}

// Load loads Hype config from hype.env.
func (c *Conf) Load() error {
	godotenv.Load("hype.env")
	if c.Editor = os.Getenv("HYPE_EDITOR"); c.Editor == "" {
		return fmt.Errorf("error, required env variable HYPE_EDITOR is unset")
	}
	if c.Port = os.Getenv("HYPE_PORT"); c.Editor == "" {
		return fmt.Errorf("error, required env variable HYPE_PORT is unset")
	}
	return nil
}
