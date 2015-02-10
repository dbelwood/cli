package cli

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"testing"
)

type configuration struct {
	ApiVersion   string `toml:"api_version"`
	ContentType  string `toml:"content_type"`
	Environments map[string]environment
}

type environment struct {
	ApiURL  string `toml:"api_url"`
	RingURL string `toml:"ring_url"`
}

type credentials struct {
	AccessToken string `toml:"access_token"`
}

const (
	configFilePath = "./_examples/test_cli.toml"
)

var (
	config configuration
)

func TestApp(t *testing.T) {
	app := NewApp()
	app.Name = "Trunk Club CLI"
	app.Flags.String(
		"environment, e",
		"development",
		"Environment in which to execute the commands",
	)

	// Load configuration from the config file
	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}
	app.Configuration = config
	app.Init = func(a *App) {
		config = a.Configuration.(configuration)
		a.Context["test_api_version"] = config.ApiVersion

	}

	// Cli Commands
	app.Commands = []Command{
		Command{
			Name:      "test",
			ShortName: "tst",
			Usage:     "Just a test command",
			Action:    func() { fmt.Println("This is a test!") },
		},
	}

	app.Run([]string{})
}
