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
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				
				err := cmd.Run()
				if err != nil {
					fmt.Println(args[0] + ": command not found")
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
		path,exist:= findPath(comnd, )
		if exist {
			fmt.Printf("%s is %s\n", comnd, path)
		} else {
			fmt.Println(comnd + ": not found")
		}

	}
}


func findPath(cmd string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	for _,dir := range strings.Split(pathEnv, ":") {
		fullPath := dir + "/" + cmd
		_,err := os.Stat(fullPath)
		if err == nil {
			return fullPath, true
		}
	// 	if isExecutable(fullPath) {
	// 		cmd := exec.Command(args[0], args[1:]...)
	// 		cmd.Stdout = os.Stdout
	// 		cmd.Stderr = os.Stderr
	// 		cmd.Stdin = os.Stdin

	// 		err := cmd.Run()
	// 		if err != nil {
	// 			fmt.Printf("%s: command not found\n", args[0])
	// }
	// 	} else {
	// 		return fullPath,true
	// 	}
		

		
	}

	return "",false
}

// func isExecutable(file string) bool {
// 	info, err := os.Stat(file)
// 	if err != nil {
// 		return false
// 	}
// 	return !info.IsDir() && info.Mode()&0111 != 0

// }