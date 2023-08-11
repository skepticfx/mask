package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary file with test data
	file, err := os.CreateTemp("", "test-config")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("foo:bar\nhello:world\n")
	if err != nil {
		t.Fatalf("Could not write to temp file: %v", err)
	}
	file.Close()

	// Load the config
	masks := loadConfig(file.Name())

	// Verify the config
	if masks["foo"] != "bar" || masks["hello"] != "world" {
		t.Errorf("Unexpected masks: %v", masks)
	}
}

func TestMaskLine(t *testing.T) {
	masks := map[string]string{
		"foo":   "bar",
		"hello": "world",
	}

	// Test case 1
	line := "foo and foo and hello"
	expected := "bar and bar and world"
	if result := maskLine(line, masks); result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test case 2
	line = "nothing to replace here"
	expected = "nothing to replace here"
	if result := maskLine(line, masks); result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
