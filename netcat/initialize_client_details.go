package netcat

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func initializeClientDetails(conn net.Conn) (*Client, error) {
    reader := bufio.NewReader(conn)

    if _, err := conn.Write([]byte(WelcomeMessage)); err != nil {
        return nil, fmt.Errorf("failed to send welcome message: %v", err)
    }

    name, err := reader.ReadString('\n')
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