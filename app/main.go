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
	"pwd" : true,
	
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

		trimmedCommand := strings.TrimSpace(command)

		if trimmedCommand == ""{
			continue
		} 

		args := strings.Split(trimmedCommand, " ")

		switch args[0] {
			case "exit" :
			if len(args) == 2 && args[1] == "0" {
				os.Exit(0)
			} 
			case "echo" :
				fmt.Println(strings.Join(args[1:], " "))

			case "type" :
				isCommand(args[1])

			case "pwd" :
				dir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
				}
				fmt.Println(dir)

			case "cd" :
				cd(args[1])
				
			default:

				_,executable := isExe(args[0])
				if executable {
					cmd := exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Stdin = os.Stdin

					err := cmd.Run()
					check(err)

				} else {
					fmt.Println(args[0] + ": command not found")
				}
				

			}


	}

}
func isCommand(comnd string) {
	if Builtin[comnd] {
		fmt.Println(comnd, "is a shell builtin")
	} else{
		path,_:= isExe(comnd)
		if  path != "" {
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
				return fullPath, true
			} else {
				return fullPath, false
			}
			
		}		
	}

	return "",false
}


func cd(dir string) {
	//dir ''
	//dir 'wrong address'
	chdir := pathBuild(dir)
	if chdir != "" {
		info,err := os.Stat(chdir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("cd: %s: No such file or directory\n", chdir)
			} else {
				fmt.Println("Error accessing path:", err)
			}
		} else if info.IsDir() {
			err = os.Chdir(chdir)
			check(err)
		}
	}
	

}

func pathBuild(uInput string) string{
	// ../ == move back once
	// ./dir == open a folder inside the current folder
	var finalPath string
	if len(uInput) > 0 {
		userParts := strings.FieldsFunc(uInput, func(r rune) bool {
			return  r == '/'
		})

		dir, _ := os.Getwd()
		dirParts := strings.FieldsFunc(dir, func(r rune) bool {
			return  r == '/'
		})
		if string(uInput[0]) == "/" {
			//return the uinput as the final path   /usr/bin
			
			return uInput
			
		} else{
			//getwd() + uinput modified
			// ./bin getwd()=> /usr + /bin[check from the last value until . show up and remove it ]

			for _,value := range userParts{
				switch value {
				case ".":
					continue
				case "..":
					getParentPath(&dirParts)
				case "~":
					homeDir := os.Getenv("HOME")
					dirParts = nil
					dirParts = append(dirParts, homeDir)

				default:
					dirParts = append(dirParts, value)
				}
			}

		}

		finalPath = "/" + strings.Join(dirParts, "/")
		
		
	}else {
		return ""
	}
	return  finalPath
}

func getParentPath(path *[]string) {
	*path = (*path)[:len(*path)-1]
}



func check(e error) {
    if e != nil {
        panic(e)
    }
}