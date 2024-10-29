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

// HandleConnection manages a single client connection
func HandleConnection(conn net.Conn, history *History) {
	defer conn.Close()

	client, err := RegisterClient(conn)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
	}

	allMessages := history.List()
	fmt.Println(allMessages)
	client.Conn.Write([]byte(allMessages))

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
		Broadcast(client, fmt.Sprintf("%s has left our chat...", client.Name), history)
	}()

	// Announce new client
	joinerMsg := FormatMessage(client.Name, "You have joined our chat...")
	client.Conn.Write([]byte(joinerMsg + "\n"))

	// displayed message on other client interfaces
	othersMsg := fmt.Sprintf("%s has joined our chat...", client.Name)
	formattedOtherMsg := FormatMessage(client.Name, othersMsg)
	history.Save(formattedOtherMsg + "\n")

	mutex.Lock()
	for otherClient := range clients {
		if otherClient != client {
			otherClient.Conn.Write([]byte(formattedOtherMsg + "\n"))
		}
	}
	mutex.Unlock()

	// Read and broadcast messages
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		message = message[:len(message)-1]
		conn.Write([]byte("\033[A\033[2K"))
		Broadcast(client, message, history)
	}
}
