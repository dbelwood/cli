package cli

import (
	"flag"
	"log"
)

// A command to run
type Command struct {
	// Command name
	Name string

	// Shorter name
	ShortName string

	// How to use this command
	Usage string

	// The function to run when this command is run
	Action func()

	// Command-specific flags.  Each Action has access to the App and Command
	// flags
	Flags flag.FlagSet

	// Sub commands, optional
	Subcommands []Command
}

// The main entry point for executing cli applications
type App struct {
	// Application name
	Name string

	// Initialization function
	Init func(*App)

	// Array of Commands, Comands can be nested
	Commands []Command

	// Configuration settings passed to each command
	Configuration interface{}

	Context map[string]interface{}

	// Flags consumed by the app
	// Applications using the cli library should define their own flags
	// after calling NewApp
	Flags *flag.FlagSet
}

func (a *App) Usage() {
}

func (a *App) Run(args []string) {
	if err := a.Flags.Parse(args); err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
	a.Init(a)
}

func NewApp() *App {
	return &App{
		Context: map[string]interface{}{},
		Flags:   flag.NewFlagSet("global", flag.ContinueOnError),
	}
}
