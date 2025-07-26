package commands

import (
	"fmt"
	"os"
)

func CommandExit(_ *Config) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
