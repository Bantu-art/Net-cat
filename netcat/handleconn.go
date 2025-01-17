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

// Broadcast sends a message to all connected clients

// HandleConnection manages a single client connection
func HandleConnection(conn net.Conn, history *History, clients map[*Client]bool, mutex *sync.Mutex) {
	defer conn.Close()

	// initialize the clients details, e.g name, will create a "Client" type
	client, err := initializeClientDetails(conn)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
	}

	// send all previous mesages to the new client
	allMessages := history.List()
	// fmt.Println(allMessages)
	client.Conn.Write([]byte(allMessages))

	fmt.Printf("\t%s, has just joined the server...\n", client.Name)

	// Add the client to the clients list
	// this is the list we will use to broadcast messages
	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	// Remove when they disconnect
	defer func() {
		mutex.Lock()
		delete(clients, client)
		mutex.Unlock()
		fmt.Printf("  %s has left the chat...\n", client.Name) // log it on server
		Broadcast(client, fmt.Sprintf("%s has left our chat...", client.Name), history, clients, mutex)
	}()

	// Announce new client
	joinerMsg := FormatMessage(client.Name, "You have joined our chat...")
	client.Conn.Write([]byte(joinerMsg + "\n"))

	// displayed message on other client interfaces
	othersMsg := fmt.Sprintf("%s has joined our chat...", client.Name)
	formattedOtherMsg := FormatMessage(client.Name, othersMsg)

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
		message = trimSpace(message)
		if len(message) == 0 || message == "\n" || message == "" || message == " " {
			conn.Write([]byte("Message cannot be empty\n"))
		} else {
			// message = message[:len(message)-1]
			conn.Write([]byte("\033[A\033[2K"))
			Broadcast(client, message, history, clients, mutex)
		}
	}
}
