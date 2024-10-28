package netcat

// import (
// 	"net"
// 	"strconv"
// 	"testing"
// )

// func TestStartServer(t *testing.T) {
// 	for i := 8000; i < 9999; i++ {
// 		port := strconv.Itoa(i)
// 		server, err := StartServer(port)
// 		if err != nil {
// 			// failed to start server on the current port: i
// 			t.Errorf("Could not start server on port %s: %v", port, err)
// 			continue // Move to the next port
// 		} else {
// 			// successfully created the server, let the tests begin...
// 			defer func(s net.Listener) {
// 				err := s.Close()
// 				if err != nil {
// 					t.Errorf("Error closing server on port %s: %v", port, err)
// 				}
// 			}(server)

// 			conn, err := net.Dial("tcp", "localhost:"+port) // try conecting to see if server is reachable
// 			if err != nil {
// 				t.Errorf("Server not reachable on port %s: %v", port, err)
// 			} else {
// 				conn.Close() // Close the connection if it succeeded
// 			}

// 			break // once successful, we can be sure server will run on a port, hence no need to continue for-loop
// 		}
// 	}
// }
