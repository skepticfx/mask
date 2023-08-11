package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
)

func main() {
	// Load the user's home directory
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Read the config file
	// Looks for ~/.mask.config by default
	configPath := path.Join(usr.HomeDir, "/.mask.config")
	masks, err := loadConfig(configPath)
	if err != nil {
		fmt.Printf("Could not find the ~/.mask.config file\n")

		// Create an example config file
		exampleContent := "word:replacement\n"
		err := os.WriteFile(configPath, []byte(exampleContent), 0644)
		if err != nil {
			fmt.Printf("Could not create the example config file: %v\n", err)
			return
		}

		fmt.Printf("An example ~/.mask.config file has been created.\n")
		masks, err = loadConfig(configPath) // Reload the newly created config
		if err != nil {
			panic(err) // Should not happen, as the file was just created
		}
	}

	// Read from stdin and mask words
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(maskLine(line, masks))
	}
}

// loadConfig reads the config file and returns a map of words to their replacements
func loadConfig(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	masks := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			masks[parts[0]] = parts[1]
		}
	}
	return masks, scanner.Err()
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
