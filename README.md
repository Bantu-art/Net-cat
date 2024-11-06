# Net-Cat Project

A TCP-based chat application implemented in Go that allows multiple clients to connect to a server and communicate in a group chat environment.

## Authors
- Anthony Oduor
- Joel Adero
- Brian Bantu

## Description

This project is a recreation of the NetCat utility with a focus on implementing a server architecture for group chat functionality. It's built using Go and implements TCP connections to facilitate communication between multiple clients through a central server.

### Features

- TCP server handling multiple client connections (1:many relationship)
- Maximum of 10 concurrent connections
- Username requirement for joining the chat
- Broadcast of messages to all connected clients
- Real-time notifications for user joins and departures
- Message history for new connections
- Timestamp and username identification for all messages
- Default port 8989 if none specified
- Linux logo welcome message

## Usage

### Starting the Server

```bash
# Default port (8989)
go run .

# Custom port
go run . 2525
```
### Connecting as a Client

```bash
nc localhost <port>
```

## Message Format
### Messages in the chat are formatted as:

```bash
[YYYY-MM-DD HH:MM:SS][username]:message
```

## Building and Running

### 1. Clone the repository

```bash
git clone https://learn.zone01kisumu.ke/git/anoduor/net-cat.git
```

### 2. Navigate to the Project directory

```bash
cd net-cat
```
### 3. Build the project

```bash
go build -o TCPChat
```
### 4. Run the server

```bash
./TCPChat [port]
```

## Error Messages

If incorrect usage: [USAGE]: ./TCPChat $port

Connection limit reached: Server notifies when maximum connections (10) is reached

## Requirements

Go 1.6 or higher

Network connectivity for client-server communication

## Testing

The project includes test files for unit testing. While inside the root folder, run the tests using:

```bash
go test ./...
```

## 🤝 Contribution

This project was developed and contributions made by:

- [Brian Bantu](https://github.com/Bantu-art)
- [Antony Oduor](https://github.com/oduortoni)
- [Joel Adero](https://github.com/Murzuqisah)


