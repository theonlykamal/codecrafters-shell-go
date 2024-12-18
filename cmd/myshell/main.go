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

func invalidCommand(command string) {
	fmt.Printf("%s: command not found\n", command)
}

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		readLine, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		if readLine == "exit 0\n" {
			os.Exit(0)
		}
		command := strings.Split(readLine, " ")[0]
		if command == "echo" {
			str := readLine[5 : len(readLine)-1]
			fmt.Println(str)
			continue
		}
		if command == "type" {
			args := strings.Split(readLine, " ")[1]
			args = args[:len(args)-1]
			if isBuiltIN(args) {
				fmt.Printf("%s is a shell builtin\n", args)
			} else {
				fmt.Printf("%s: not found\n", command)
			}

			continue
		}
		invalidCommand(command)
	}

}
