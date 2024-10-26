package netcat

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

const WelcomeMessage = `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 __| ".        |\dS"qML
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
[ENTER YOUR NAME]: `

type Client struct {
	Conn net.Conn
	Name string
}

var (
	clients = make(map[*Client]bool) // Using a map as a set to track clients
	mutex   = &sync.Mutex{}          // To safely modify the clients map
)

// Broadcast sends a message to all connected clients
func Broadcast(message string) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		client.Conn.Write([]byte(message + "\n"))
	}
}

// HandleConnection manages a single client connection
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	client, err := RegisterClient(conn)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
	}

	fmt.Printf("New client: (%s) registered\n", client.Name)
	// Add to clients list
	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	// Remove when they disconnect
	defer func() {
		mutex.Lock()
		delete(clients, client)
		mutex.Unlock()
		Broadcast(fmt.Sprintf("%s has left our chat...", client.Name))
	}()

	// Announce new client
	Broadcast(fmt.Sprintf("%s has joined our chat...", client.Name))

	// Read and broadcast messages
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fullMessage := fmt.Sprintf("[%s]: %s", client.Name, message)
		Broadcast(fullMessage)
	}
}
