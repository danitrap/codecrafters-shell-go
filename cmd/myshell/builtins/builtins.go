package builtins

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/helpers"
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
	_, err := GetBuiltin(cmd)
	if err == nil {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	if p, err := helpers.GetExecutable(cmd); err == nil {
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

func Cd(args []string) {
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", args[1])
	}
}

func GetBuiltin(cmd string) (BuiltinCommand, error) {
	builtin, ok := builtins[cmd]
	if !ok {
		return BuiltinCommand{}, fmt.Errorf("%s: not a builtin command", cmd)
	}
	return builtin, nil
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
		"cd": {
			Name:           "cd",
			Implementation: Cd,
		},
	}
}
