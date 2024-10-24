package netcat

import "net"

func StartServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}
