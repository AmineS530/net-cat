package main

import (
	"os"
)

func SaveToFile(name, message string) {
	file, _ := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer file.Close()
	file.WriteString(message)
}

func prevMessage() string {
	data, _ := os.ReadFile("txtFiles/messageHistory"+"["+port+"]"+".txt")
	return string(data)
}

func IsNameTaken(name string) bool {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for _, client := range clients {
		if client.Name == name {
			return true // Name is taken
		}
	}
	return false
}

func IsPrint(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(rune(str[i]) >= 32 && rune(str[i]) <= 126 || str[i] == 10) {
			return false
		}
	}
	return true
}
