package commands

import (
	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
	"github.com/RyutoLBX/pokedexcli/internal/pokecache"
)

type Config struct {
	Pokedex      pokeapi.Pokedex
	Next         *string
	Previous     *string
	MapCache     *pokecache.Cache
	ExploreCache *pokecache.Cache
	PokemonCache *pokecache.Cache
}
