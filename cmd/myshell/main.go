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

func IsBuiltIN(command string) bool {
	if command == "echo" || command == "exit" || command == "type" {
		return true
	}
	return false

}

func WhereIs(args string, envVarString string) bool {
	//Not a Built In coommand
	PATH := strings.Split(os.Getenv(envVarString), ":")
	//All paths
	for _, path := range PATH {
		contents, _ := os.ReadDir(path)

		//All files
		for _, file := range contents {
			if !file.IsDir() && file.Name() == args {
				fmt.Printf("%s is %s/%s\n", args, path, args)
				return true
			}
		}
	}
	return false
}

func ExecCommand(command string, args ...string) (bool, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(string(output))
	return true, err
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

			if IsBuiltIN(args) {
				fmt.Printf("%s is a shell builtin\n", args)
			} else {
				if !WhereIs(args, "PATH") {
					fmt.Printf("%s: not found\n", args)
				}
			}
			continue
		}

		success, _ := ExecCommand(command, readLine[1:])
		if success {
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}

}
