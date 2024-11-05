package netcat

import "fmt"

func Broadcast(sender *Client, message string, history *History) {
	formattedMsg := FormatMessage(sender.Name, message)
	history.Save(formattedMsg + "\n")
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Fprintln(sender.Conn, formattedMsg)

	for client := range clients {
		if client == sender {
			continue
		}
		if(client != nil) {
			client.Conn.Write([]byte(formattedMsg + "\n"))
		}
	}
}
