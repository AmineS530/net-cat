package netcat

import (
	"os"
	"time"
)

// Clears the previous message file by truncating it
func ClearMessageHistory() {
	file, _ := os.OpenFile("txtFiles/messageHistory_["+port+"].txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	file.Truncate(0)
}

// Saves the start of a new chat log with the current time
func SaveNewChatLog() {
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	SaveToFile("txtFiles/logs.txt", "------------------------new chat started at ["+formattedTime+"]--------------------------------\n\n\n")
}

func SaveToFile(name, message string) {
	file, _ := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer file.Close()
	file.WriteString(message)
}

func prevMessage() string {
	data, _ := os.ReadFile("txtFiles/messageHistory_[" + port + "].txt")
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
