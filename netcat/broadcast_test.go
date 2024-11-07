package netcat

import (
	"sync"
	"testing"
)

func TestBroadcast(t *testing.T) {
	clients := make(map[*Client]bool)
	history := NewHistory()
	mutex := &sync.Mutex{}

	conn1 := &MockTcpConn{
		OutputBuffer: make([]byte, 30),
	}
	client1 := &Client{
		Conn: conn1,
		Name: "c_one",
	}
	conn2 := &MockTcpConn{
		OutputBuffer: make([]byte, 30),
	}
	client2 := &Client{
		Conn: conn2,
		Name: "c_two",
	}
	conn3 := &MockTcpConn{
		OutputBuffer: make([]byte, 30),
	}
	client3 := &Client{
		Conn: conn3,
		Name: "c_tre",
	}
	clients[client1] = true
	clients[client2] = true
	clients[client3] = true

	Broadcast(client1, "Hello", history, clients, mutex)
	Broadcast(client1, "world", history, clients, mutex)
	Broadcast(client1, "kenya", history, clients, mutex)

	for client := range clients {
		if client.Conn == nil {
			t.Errorf("Empty client")
		}
	}
	list := history.List()
	if list == "" {
		t.Errorf("History can not be empty after save operations\n")
	}
}
