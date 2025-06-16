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


	for i := 1; i > 0; i++ {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			panic(err)
		}

		trimmedCommand := command[:len(command)-1]
		if trimmedCommand == "exit 0" {
			os.Exit(0)
		}

		com := strings.Split(trimmedCommand, " ")

		if com[0] == "echo" {
			for _,v := range com[1:] {
				fmt.Print(v, " ")
			}
			fmt.Println()
			continue
		}


		//Since the string returned by ReadString('\n') includes a trailing newline, use command[:len(command)-1] to remove it.
		fmt.Println(trimmedCommand + ": command not found")
	}

}
