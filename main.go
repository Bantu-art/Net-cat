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

	listener, err := netcat.StartServer(port)
	if err != nil {
		fmt.Printf("\n[SERVER STARTING ERROR] Failed to start server.\nMessage: %v.\nIs the port already in use?\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("\n\t--------------------------------------\n\n\t  Server listening on the port :%s  \n\n\t--------------------------------------\n\n", port)

	history := netcat.NewHistory()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("\t[ERROR] Failed to accept connection: %v\n", err)
			continue
		} else {
			nClients := len(clients)
			if nClients >= MAXCLIENTS {
				conn.Write([]byte("[ERROR] server is not accepting any more connections. Try again later.... \n\tPress ENTER to exit..."))
				conn.Close()
			} else {
				go netcat.HandleConnection(conn, history, clients, mutex)
			}
		}
	}
}
