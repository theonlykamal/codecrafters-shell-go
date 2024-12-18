package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func isBuiltIN(command string) bool {
	if command == "echo" || command == "exit" || command == "type" {
		return true
	}
	return false

}

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		readLine, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}

		readLine = readLine[:len(readLine)-1]

		if readLine == "exit 0" {
			os.Exit(0)
		}
		command := strings.Split(readLine, " ")[0]

		if command == "echo" {
			str := readLine[5:]
			fmt.Println(str)
			continue
		}
		if command == "type" {
			args := strings.Split(readLine, " ")[1]

			if isBuiltIN(args) {
				fmt.Printf("%s is a shell builtin\n", args)
			} else {
				//Not a Built In coommand
				PATH := strings.Split(os.Getenv("PATH"), ":")
				found := false
				//All paths
				for _, path := range PATH {
					contents, err := os.ReadDir(path)
					if err != nil {
						return
					}

					//All files
					for _, file := range contents {
						if !file.IsDir() && file.Name() == args {
							fmt.Printf("%s is %s/%s\n", args, path, args)
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				continue

			}
			fmt.Printf("%s: not found\n", args)

			continue
		}
		fmt.Printf("%s: command not found\n", command)
	}

}
