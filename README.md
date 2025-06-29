# Local Chat Messenger (Implmeneted in Go)
A simple UDP-based client-server chat application built with Go. This project demonstrates UDP socket programming, concurrent connections, and bidirectional communication.

## Features
- **UDP Socket Communication**: Connectionless messaging between clients and server
- **Multiple Concurrent Clients**: Server can handle multiple clients simultaneously
- **Bidirectional Chat**: Clients can send messages and receive responses from the server
- **Automatic Port Selection**: Client ports are automatically assigned by the OS
- **Echo Server**: Server responds with echoed messages plus confirmation text

## Server & Client Description
### Server
- Listens on `127.0.0.1:8090` for incoming UDP messages
- Receives messages from any number of clients
- Responds with echoed message plus "(echoing from server)" suffix
- Displays client IP and port for each received message

### Client
- Connects to server at `127.0.0.1:8090`
- Uses automatically assigned local port (via `nil` local address)
- Sends user input to server and displays server responses
- Supports graceful exit with "exit" command

## Getting Started
### Prerequisites
Go 1.16 or higher

### Running the Application
1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd local_chat_messenger
   ```

2. **Start the server** (in first terminal)
   ```bash
   go run main.go server.go client.go server
   ```

3. **Start client(s)** (in separate terminal(s))
   ```bash
   go run main.go server.go client.go client
   ```

4. **Chat away!**
   - Type messages in the client terminal
   - See echoed responses from the server
   - Type "exit" to close connections

### Testing Multiple Clients
Open multiple terminals and run the client command in each. All clients can communicate with the server simultaneously.

## Example Usage
**Server Output:**
```
127.0.0.1(54321): hello server
127.0.0.1(54322): how are you?
127.0.0.1(54321): i'm doing great
```

**Client Output:**
```
Type your message: hello server
127.0.0.1(8090): hello server (echoing from server)
Type your message: how are you?
127.0.0.1(8090): how are you? (echoing from server)
```
