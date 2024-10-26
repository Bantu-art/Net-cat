package netcat

func Broadcast(name, message string) {
	formattedMsg := FormatMessage(name, message)
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		client.Conn.Write([]byte(formattedMsg + "\n"))
	}
}
