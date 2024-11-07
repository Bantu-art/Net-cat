package main

import (
	"fmt"

	"netcat/netcat"
)

func main() {
	port := netcat.GetPort()
	if port == "" {
		return
	}

	listener, err := netcat.StartServer(port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on the port :%s\n", port)

	history := netcat.NewHistory()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go netcat.HandleConnection(conn, history)
	}
}
