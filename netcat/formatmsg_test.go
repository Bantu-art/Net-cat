package netcat

import (
	"strings"
	"testing"
)

func TestFormatmsg(t *testing.T) {
	table := []struct {
		Name    string
		Message string
	}{
		{
			Name:    "toni",
			Message: "Hello there",
		},
		{
			Name:    "Alexa",
			Message: "Its 123 note",
		},
	}

	for _, item := range table {
		formatted := FormatMessage(item.Name, item.Message)

		if !strings.Contains(formatted, item.Name) {
			t.Errorf("Expected %q to contain  the name, %q, but it does not\n", formatted, item.Name)
		}

		if !strings.Contains(formatted, item.Message) {
			t.Errorf("Expected %q to contain the message, %q, but it does not\n", formatted, item.Message)
		}
	}
}
