package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BuiltinCommand struct {
	Name           string
	Implementation func(args []string)
}

var builtins map[string]BuiltinCommand

func Exit(args []string) {
	os.Exit(0)
}

func Echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func Type(args []string) {
	cmd := args[1]
	if _, ok := builtins[cmd]; ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	fmt.Printf("%s: not found\n", cmd)
}

func init() {
	builtins = map[string]BuiltinCommand{
		"exit": {
			Name:           "exit",
			Implementation: Exit,
		},
		"echo": {
			Name:           "echo",
			Implementation: Echo,
		},
		"type": {
			Name:           "type",
			Implementation: Type,
		},
	}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		parts := strings.Split(strings.TrimSpace(input), " ")

		if builtin, ok := builtins[parts[0]]; ok {
			builtin.Implementation(parts)
			continue
		}

		fmt.Println(parts[0] + ": command not found")
	}
}
