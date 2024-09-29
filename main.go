package main

import (
	"fmt"
	"net"

	netcat "main/goFiles"
)

func init() {
	netcat.ClearMessageHistory()
	netcat.SaveNewChatLog()
}

func main() {
	// Start the TCP server and handle errors
	listener, err := netcat.StartTCPServer()
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return // Exit if the server couldn't start
	}
	defer listener.Close() // Close listener when the program exits

	handleCommands(listener)
}

// Handles user input commands in a loop
func handleCommands(listener net.Listener) {
	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "list":
			netcat.ListClients()
		case "kick":
			netcat.KickClient()
		case "kill":
			netcat.KillServer(listener)
			return
		}
	}
}
