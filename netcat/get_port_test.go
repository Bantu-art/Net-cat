package netcat

import (
	"os"
	"testing"
)

func TestGetport(t *testing.T) {
	table := []struct {
		Args   []string
		Port string
	}{
		{[]string{"./TCPChat"}, "8989"},              
		{[]string{"./TCPChat", "8080"}, "8080"},
		{[]string{"./TCPChat", "8080", "extraargument"}, ""},
	}

	for _, item := range table {
		origArgs := os.Args // Save original os.Args
		defer func() { os.Args = origArgs }() // after test, restore the os.Args

		os.Args = item.Args
		port := GetPort()

		if port != item.Port {
			t.Errorf("Given args as, %v, expected %v, got %v", item.Args, item.Port, port)
		}
	}
}
