package commands

import (
	"github.com/RyutoLBX/pokedexcli/internal/pokecache"
)

type Config struct {
	Next         *string
	Previous     *string
	MapCache     *pokecache.Cache
	ExploreCache *pokecache.Cache
}
