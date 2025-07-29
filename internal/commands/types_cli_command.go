package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}
