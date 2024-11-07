package netcat

import (
	"fmt"
	"os"
)

func GetPort() string {
	const defaultPort = "8989"

	switch len(os.Args) {
	case 1:
		return defaultPort
	case 2:
		return os.Args[1]
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return ""
	}
}
