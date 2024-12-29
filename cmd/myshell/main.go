package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
	if IsBuiltin(cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	if p, err := GetExecutable(cmd); err == nil {
		fmt.Printf("%s is %s/%s\n", cmd, p, cmd)
		return
	}
	fmt.Printf("%s: not found\n", cmd)
}

func Pwd(args []string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pwd)
}

func IsBuiltin(cmd string) bool {
	_, ok := builtins[cmd]
	return ok
}

func GetExecutable(cmd string) (string, error) {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	for _, p := range paths {
		if file, err := os.Stat(p + "/" + cmd); err == nil {
			if file.Mode().IsRegular() && file.Mode()&0111 != 0 {
				return p, nil
			}
		}
	}
	return "", fmt.Errorf("%s: not found", cmd)
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
		"pwd": {
			Name:           "pwd",
			Implementation: Pwd,
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

		if p, err := GetExecutable(parts[0]); err == nil {
			output, err := exec.Command(p+"/"+parts[0], parts[1:]...).Output()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Fprint(os.Stdout, string(output))
			continue
		}

		fmt.Println(parts[0] + ": command not found")
	}
}
