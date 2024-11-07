package netcat

import (
	"fmt"
	"time"
)

func FormatMessage(name, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if name == "" {
		return fmt.Sprintf("[%s] %s", timestamp, message)
	}

	// Regular chat message
	return fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)
}
