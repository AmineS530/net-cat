package netcat

import (
	"fmt"
	"os"
)

var port = GetPort()

func PortChecker(port string) bool {
	if len(port) > 5 || len(port) < 4 {
		return false
	}
	for _, c := range port {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

func GetPort() (port string) {
	port = ":8989"
	args := os.Args
	if len(args) == 2 {
		if PortChecker(args[1]) {
			port = ":" + args[1]
		} else {
			fmt.Println("Please Check your Port Nubmer is not correct\nEx: go run . 8989")
			os.Exit(0)
		}
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	return
}
