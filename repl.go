package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))

}
