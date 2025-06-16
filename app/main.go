package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	builtin := []string{"echo","exit"}
	for i := 1; i > 0; i++ {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			panic(err)
		}

		trimmedCommand := command[:len(command)-1]

		args := strings.Split(trimmedCommand, " ")
		
		switch args[0] {
			case "exit" :
			if len(args) == 2 && args[1] == "0" {
				os.Exit(0)
			} 
			case "echo" :
				fmt.Println(strings.Join(args[1:], " "))
		
			case "type" :
				for _,isbuitlin := range builtin {
					if len(args) == 2 && isbuitlin == args[1] {
						fmt.Println(args[1], "is a shell builtin")
					}
				}
			default:
				fmt.Println(args[0] + ": command not found")
		}


	}

}
