package commands

import (
	"fmt"
	"os"
)

// Exits the REPL loop
func CommandExit(_ *Config) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
