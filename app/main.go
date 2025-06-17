package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var Builtin = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	
}
func main() {

	
	for  {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		

		if err != nil {
			panic(err)
		}

		trimmedCommand := command[:len(command)-1]

		args := strings.Split(trimmedCommand, " ")
		if len(args) > 1 {
			switch args[0] {
			case "exit" :
			if len(args) == 2 && args[1] == "0" {
				os.Exit(0)
			} 
			case "echo" :
				fmt.Println(strings.Join(args[1:], " "))
		
			case "type" :
				isCommand(args[1])
				

			default:

				_,executable := isExe(args[0])
				if executable {
					cmd := exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Stdin = os.Stdin
					
					err := cmd.Run()
					if err != nil {
						fmt.Println(args[0] + ": command not found")
					}
				}
				
			}
		} else if len(args) == 1 {
			fmt.Println(args[0] + ": command not found")
		}


	}

}
func isCommand(comnd string) {
	if Builtin[comnd] {
		fmt.Println(comnd, "is a shell builtin")
	} else{
		path,exist:= isExe(comnd, )
		if exist {
			fmt.Printf("%s is %s\n", comnd, path)
		} else {
			fmt.Println(comnd + ": not found")
		}

	}
}


func isExe(cmd string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	for _,dir := range strings.Split(pathEnv, ":") {
		fullPath := dir + "/" + cmd
		info,err := os.Stat(fullPath)
		if err == nil {
			if !info.IsDir() && info.Mode()&0111 != 0 {
				return "", true
			}
			return fullPath, true
		}		
	}

	return "",false
}

