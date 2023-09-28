package main
import "fmt"


func commandHelp(cfg *config, args ...string) error{
	fmt.Println("Welcome to Pokedex")
			fmt.Println("Here are the few commands in help")
			availableCommands:=getCommands()
			for _,cmd := range availableCommands {
				fmt.Println(cmd.name)
				fmt.Println(cmd.description)
			}
			fmt.Println("")
			return nil
}