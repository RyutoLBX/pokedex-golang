package pokeapi

import "fmt"

func NewPokedex() Pokedex {
	return Pokedex{Pokemons: make(map[string]Pokemon)}
}

func (p Pokedex) AddToPokedex(pokemon Pokemon) {
	_, exists := p.Pokemons[pokemon.Name]
	if !exists {
		p.Pokemons[pokemon.Name] = pokemon
		fmt.Printf("%s added to the pokedex.\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s already recorded in the pokedex!\n", pokemon.Name)
	}
}
