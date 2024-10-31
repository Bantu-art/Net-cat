package netcat

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

/*
* initializeClientDetails
* Takes a net.Conn and tries to create a Client object with the details of a client
* Details might include the name, and other details to easily identify a client
* It accomplishes all these by promting the user through the Conncetion object
* returns a pointer to  Client and an error
*
*/
func initializeClientDetails(conn net.Conn) (*Client, error) {
    reader := bufio.NewReader(conn)

    if _, err := conn.Write([]byte(WelcomeMessage)); err != nil {
        return nil, fmt.Errorf("failed to send welcome message: %v", err)
    }

    name, err := reader.ReadString('\n') // fetch client name
    if err != nil {
        return nil, fmt.Errorf("failed to read client name: %v", err)
    }

    name = strings.TrimSpace(name)
    if name == "" {
        conn.Write([]byte("Name cannot be empty\n"))
        return nil, fmt.Errorf("empty name provided")
    }

    return &Client{
        Conn: conn,
        Name: name,
    }, nil
}