package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var port = GetPort()

func main() {
	// turncate the prev message file
	file, _ := os.OpenFile("txtFiles/messageHistory.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	file.Truncate(0)
	///save start of new chat with time
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	SaveToFile("txtFiles/logs.txt", "------------------------new chat started at ["+formattedTime+"]--------------------------------\n\n\n")

	// Start TCP server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on the Port: ", port[1:])

	// Accept connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleClient(conn)
		}
	}()

	// Periodically list connected clients (optional)
	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "list":
			listClients()
		case "kick":
			kickClient()
		}
	}
}
