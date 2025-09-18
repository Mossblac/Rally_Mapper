package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Mossblac/Rally_Mapper/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Rally_Mapper > ")
		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)
		var cInput string

		if len(words) == 1 {
			cInput = words[0]

		} else {
			cInput = words[0]

		}
		if len(words) > 2 {
			fmt.Printf("invalid input: single word only, use dashes for areas\n")
		} else {
			_, exists := commands.CommandMap[cInput]
			if exists {
				c := commands.CommandMap[cInput]
				err := c.Callback(cInput)
				if err != nil {
					fmt.Printf("error: %v", err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}

}
