package netcat

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("New connection from: %s\n", conn.RemoteAddr().String())
}
