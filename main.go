package main

import (
	"fmt"
	"sync"

	"netcat/netcat"
)

var MAXCLIENTS = 2

var (
	clients = make(map[*netcat.Client]bool) // Using a map as a set to track clients
	mutex   = &sync.Mutex{}                 // To safely modify the clients map
)

func main() {
	port := netcat.GetPort()
	if port == "" {
		return
	}

	fmt.Println("MAPLEN: ", len(clients))

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
		} else {
			nClients := len(clients)
			fmt.Println("Number is ", nClients)
			if nClients >= MAXCLIENTS {
				conn.Write([]byte("Not accepting any more connections"))
				conn.Close()
			} else {
				go netcat.HandleConnection(conn, history, clients, mutex)
			}
		}
	}
}
