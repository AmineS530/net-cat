package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	if len(clients) == 10 {
		conn.Write([]byte("Connection is Full in the server"))
		return
	}
	clientAddr := conn.RemoteAddr().String()
	// linuxLogo
	linuxLogo, _ := os.ReadFile("txtFiles/linuxLogo.txt")
	conn.Write(linuxLogo)

	for {
		// read name
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		reader := bufio.NewReader(conn)
		name, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		name = strings.TrimSpace(name)
		if IsNameTaken(name) {
			conn.Write([]byte("Name already taken, please choose another.\n"))
		} else if len(name) > 0 && IsPrint(name) && !IsNameTaken(name) {
			clientsMutex.Lock()
			clients[clientAddr] = Client{Conn: conn, Name: name}
			clientsMutex.Unlock()
			break
		} else {
			conn.Write([]byte("Please enter a valid name.\n"))
		}
	}
	reader := bufio.NewReader(conn)

	writeToClients(" has joined the chat...\n", clientAddr, false)

	clients[clientAddr].Conn.Write([]byte(prevMessage()))

	bl := true
	for {
		// check for empty message
		if bl {
			Status()
		}
		bl = true
		// Read data until newline or EOF (you can modify the delimiter if needed)
		message, err := reader.ReadString('\n')
		if err != nil {
			writeToClients(" has left the chat...\n", clientAddr, false)
			Status()
			delete(clients, clientAddr)
			break // Exit loop if the client disconnects
		}

		if len(message) == 1 {
			conn.Write([]byte("Empty message was not sent\n" + geneateMessage(clients[clientAddr].Name)))
			bl = false
		} else {
			clientsMutex.Lock()
			lastSentTime, ok := lastMessageTime[clientAddr]
			if !ok || time.Since(lastSentTime) >= cooldownTime {
				// Update last message time
				lastMessageTime[clientAddr] = time.Now()
				if IsPrint(message) && !strings.HasPrefix(message, "\033[") && !(len(strings.TrimSpace(message)) == 0) {
					writeToClients(message, clientAddr, true)
				} else {
					conn.Write([]byte("Invalid input. Please try again.\n" +
						geneateMessage(clients[clientAddr].Name)))
					bl = false
				}
			} else {
				conn.Write([]byte("You are sending messages too quickly. Please wait a moment.\n" +
					geneateMessage(clients[clientAddr].Name)))
				bl = false
			}
			clientsMutex.Unlock()
		}
	}
	// Remove client from the map when they disconnect
	clientsMutex.Lock()
	delete(clients, clientAddr)
	clientsMutex.Unlock()
}
