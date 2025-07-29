package commands

type cliCommand struct {
	name        string
	params      []string
	description string
	Callback    func(*Config, []string) error
}
