package main

import (
	"time"

	"github.com/elaranel/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	prm := &parameters{}

	startRepl(cfg, prm)
}
