package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		if (parts[0]) == "exit" {
			os.Exit(0)
		}

		fmt.Println(parts[0] + ": command not found")
	}
}
