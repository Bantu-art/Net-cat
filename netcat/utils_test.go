package netcat

import (
	"fmt"
	"testing"
)

func TestReplaceSpecialCharacters(t *testing.T) {
	tTests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "String with newline in the middle",
			input:    "Hello\\nWorld",
			expected: "Hello World",
		},
		{
			name:     "String without any special characters",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "String with a tab and an bell character",
			input:    "Hello\\tWorld\\a",
			expected: "Hello World ",
		},
	}
	for _, tItem := range tTests {
		t.Run(tItem.name, func(t *testing.T) {
			result := replaceSpecialcharacters(tItem.expected)
			fmt.Printf("Result: %q\nExpected: %q\n", result, tItem.input)
			if result != tItem.expected {
				t.Errorf("replaceSpecialcharacters() args = %v, want %v", result, tItem.expected)
			}
		})
	}
}
