package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, args ...string) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", pokemonResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
