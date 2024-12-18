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

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		if command == "exit 0\n" {
			os.Exit(0)
		}
		if strings.Split(command, " ")[0] == "echo" {
			str := command[5 : len(command)-1]
			fmt.Println(str)
		}
		fmt.Printf("%s: command not found\n", command[:len(command)-1])
	}

}
