package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtins"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/helpers"
)

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

		builtin, err := builtins.GetBuiltin(parts[0])
		if err == nil {
			builtin.Implementation(parts)
			continue
		}

		if p, err := helpers.GetExecutable(parts[0]); err == nil {
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
