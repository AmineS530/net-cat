package netcat

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Client struct {
	Name string
	Conn net.Conn
}

var (
	clients         = make(map[string]Client)    // Map to store connected clients
	clientsMutex    sync.Mutex                   // Mutex to synchronize access to clients map
	cooldownTime    = 1500 * time.Millisecond    // Cooldown duration
	lastMessageTime = make(map[string]time.Time) // Track last message time per client
	now             = time.Now()
	formattedTime   = now.Format("2006-01-02 15:04:05")
)

func geneateMessage(name string) string {
	return "[" + formattedTime + "]" + "[" + name + "]" + ":"
}

func writeToClients(message string, clientAddr string, flag bool) {
	if flag {
		message = "\n" + geneateMessage(clients[clientAddr].Name) + message
		SaveToFile("txtFiles/messageHistory.txt", message[1:])
	} else {
		message = "\n" + clients[clientAddr].Name + message
		SaveToFile("txtFiles/logs.txt", geneateMessage("Client Name: "+clients[clientAddr].Name+" || Client Adress "+clientAddr)+message[1:])
	}

	loopAll(message, clientAddr)
}

func Status() {
	for i, j := range clients {
		if j.Name != "" {
			j.Conn.Write([]byte(geneateMessage(clients[i].Name)))
		}
	}
}

// Function to handle each client connection

func loopAll(message, clientAddr string) {
	for i, j := range clients {
		if i != clientAddr {
			if j.Name != "" && message != "" {
				j.Conn.Write([]byte(message))
			}
		}
	}
}

// Function to display all connected clients
func ListClients() {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	fmt.Println("Connected clients:")
	id := 0
	for addr := range clients {
		fmt.Printf("[%d] username: [%s] | Client Address: [%s] \n", +id, clients[addr].Name, addr)
		id++
	}
}

func KickClient() {
	ListClients()
	fmt.Print("Enter the address of the client to kick: ")
	var addrToKick string
	fmt.Scanln(&addrToKick)
	for addr, client := range clients {
		if addr == addrToKick {
			client.Conn.Write([]byte("You have been kicked from the chat.\n"))
			break
		}
	}

	if addrToKick != "" {
		// Close the client's connection
		kickedUser := clients[addrToKick].Name
		clients[addrToKick].Conn.Close()
		delete(clients, addrToKick)
		writeToClients(fmt.Sprintf("%s was kicked from the chat...\n", kickedUser), "", false)
		fmt.Printf("Kicked client: %s\n", kickedUser)
	} else {
		fmt.Println("Client not found.")
	}
}

func KillServer(listener net.Listener) {
	fmt.Println("Shutting down server in 10 seconds")
	writeToClients("Server will shut down in 10 seconds...\n", "", false)
	time.Sleep(10 * time.Second)
	listener.Close()
}
