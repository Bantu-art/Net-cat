package netcat

import (
	"testing"
)

/*
* initializeClientDetails function takes a net.Conn and prompts the client that is represented by it to enter:
* 1) name
* It then creates a Client object and returns it
* We need to mock a net.Conn object such that we have control over the buffer that is being written to by this function
* With this ability, we can check whether the initialized Client object has the name we expect
 */
func TestInitializeClientDetails(t *testing.T) {
	name := "Antony"
	conn := &MockTcpConn{
		InputBuffer: []byte(name + "\n"),
	}

	// Call the function to test
	// it will read data from the InputBuffer to create the Client eventually
	// this name will eventually be tested against the name
	client, err := initializeClientDetails(conn)
	if err != nil {
		t.Errorf("Error initializing client details: %v", err)
		return
	}

	if client.Name != name {
		t.Errorf("Expected client name to be %s, but got %s", name, client.Name)
	}
}
