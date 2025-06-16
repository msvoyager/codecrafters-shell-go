package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var Builtin = []string{"echo","exit", "type"}
func main() {

	
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
				out := isBuiltIn(args)
				if out {
					fmt.Println(args[1], "is a shell builtin")
				} else {
					fmt.Println(args[1] + ": not found")
				}
				

			default:
				fmt.Println(args[0] + ": command not found")
		}


	}

}

func isBuiltIn(args []string) bool{
	for _,v := range Builtin {
		if len(args) == 2 && v == args[1] {
			
			return true
		}
						
				
	}

	return false
}