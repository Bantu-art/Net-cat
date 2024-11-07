package netcat

import "strings"

type History struct {
    Messages []string
}

func NewHistory() *History {
	return &History{
		Messages: []string{},
	}
}

func (h *History)Save(msg string) {
	h.Messages = append(h.Messages, msg)
}

func (h History)List() string {
	msgs := strings.Join(h.Messages, "")
	return msgs
}
