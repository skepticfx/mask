package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	// Read masks from config file
	masks := loadConfig(defaultConfigPath())

	// Read from stdin, mask words and write to stdout and stderr
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	for stdin.Scan() {
		if _, err := stdout.WriteString(maskLine(stdin.Text(), masks) + "\n"); err != nil {
			bail(err.Error())
		}
	}

	if err := stdout.Flush(); err != nil {
		bail(err.Error())
	}
}

// maskLine replaces substrings in the line with their masks from the config
func maskLine(line string, masks map[string]string) string {
	// Ensure everything is lowercased
	line = strings.ToLower(line)

	// Iterate through the masks and replace all instances of each key with its value
	for word, replacement := range masks {
		line = strings.ReplaceAll(line, word, strings.ToLower(replacement))
	}
	return line
}
