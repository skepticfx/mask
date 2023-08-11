package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
)

// Masks holds the config for substring matches and their replacements
type Masks map[string]string

// loadConfig reads the config file and returns a map of words to their replacements
func loadConfig(path string) Masks {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Config file %s does not exist.", path)
		// init an example config that can be loaded
		initConfig(path)
	}

	file, err := os.Open(path)
	if err != nil {
		bail(err.Error())
	}
	defer file.Close()

	masks := make(Masks)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			masks[parts[0]] = parts[1]
		}
	}
	return masks
}

// Looks for ~/.mask.config by default
func defaultConfigPath() string {
	// Load the user's home directory
	usr, err := user.Current()
	if err != nil {
		bail(err.Error())
	}
	configPath := path.Join(usr.HomeDir, "/.mask.config")
	return configPath
}

func initConfig(path string) {
	// Create an example config file
	exampleContent := "word:replacement\n"
	err := os.WriteFile(path, []byte(exampleContent), 0644)
	if err != nil {
		bail(fmt.Sprintf("Could not create the example config file: %v", err))
	}
	fmt.Fprintln(os.Stdout, "An example ~/.mask.config file has been created.")
}
