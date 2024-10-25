package netcat

import (
	"fmt"
	"net"
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

// HandleConnection manages a single client connection
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	client, err := RegisterClient(conn)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
	}

	fmt.Printf("New client registered: %s\n", client.Name)
}
