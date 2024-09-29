package netcat

import (
	"fmt"
	"net"
)

// Starts the TCP server and listens for incoming connections
func StartTCPServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err // Return the error to be handled in main
	}

	fmt.Println("Listening on the Port: ", port[1:])

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
					// Listener closed, exit the loop
					fmt.Println("Listener closed, exiting connection loop")
					return
				}
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleClient(conn)
		}
	}()

	return listener, nil // Return the listener if everything is okay
}
