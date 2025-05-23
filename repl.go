package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/elaranel/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config, prm *parameters) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type parameters struct {
	Key string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show pokedex",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
