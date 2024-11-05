package netcat

import (
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
			if result != tItem.expected {
				t.Errorf("replaceSpecialcharacters() args = %v, want %v", result, tItem.expected)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	tTests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "String with newline at the end",
			input:    "Hello World\n",
			expected: "Hello World",
		},
		{
			name:     "Ordinary string",
			input:    "Hello World\t",
			expected: "Hello World",
		},
		{
			name:     "String with a tab and an bell character",
			input:    "Hello World\\a",
			expected: "Hello World",
		},
	}
	for _, tItem := range tTests {
		t.Run(tItem.name, func(t *testing.T) {
			result := trimSpace(tItem.input)
			if result != tItem.expected {
				t.Errorf("replaceSpecialcharacters() args = %v, want %v", result, tItem.expected)
			}
		})
	}
}
