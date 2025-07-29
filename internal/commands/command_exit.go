package commands

import (
	"fmt"
	"os"
)

// CommandExit exits the REPL loop.
func CommandExit(_ *Config, _ []string) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
