package netcat

import (
	"strings"
	"testing"
)

func TestHistorySave(t *testing.T) {
	h := NewHistory()

	messages := []string{
		"Hello there",
		"I am a user",
		"This is a history test",
	}

	for _, message := range messages {
		h.Save(message)
	}

	list := h.List()
	joinedMsgs := strings.Join(messages, "")

	if joinedMsgs != list {
		t.Errorf("Expected the list to be similar to the joined list")
	}

	joinedMsgs = strings.Join(messages, "+")
	if joinedMsgs == list {
		t.Errorf("Expected the list to be different to the joined list")
	}
}

