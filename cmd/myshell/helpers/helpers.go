package helpers

import (
	"fmt"
	"os"
	"strings"
)

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
