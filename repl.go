package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanned := cleanInput(text)
		if len(cleanned) == 0 {
			continue
		}
		commandName:= cleanned[0]
		args := []string{}
		if len(cleanned) > 1{
			args = cleanned[1:]

		}
		availableCommands:=getCommands()
		command,ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err:= command.callback(cfg, args...)
		if  err != nil {
			fmt.Println(err)
			
		}

		

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit":{
			name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
		},
		"map":{
			name: "map",
			description: "Lists some location areas",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Lists the previous page of location areas",
			callback: commandMapb,
		},
		"explore":{
			name: "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback: commandExplore,
		},
		"catch":{
			name: "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect {pokemon_name}",
			description: "View inforamtion about a caught pokemon",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "View all the pokemon in your pokedex",
			callback: commandPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
