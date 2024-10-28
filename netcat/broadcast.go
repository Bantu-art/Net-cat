package netcat

func Broadcast(name, message string, history *History) {
	formattedMsg := FormatMessage(name, message)
	history.Save(formattedMsg)
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		client.Conn.Write([]byte(formattedMsg + "\n"))
	}
}
