package main

import "os"

func SaveToFile(name, message string) {
	file, _ := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer file.Close()
	file.WriteString(message)
}

func prevMessage() string {
	data, _ := os.ReadFile("prevMessages.txt")
	return string(data)
}

func isValidName(name string) bool {
	for i := 0; i < len(name); i++ {
		if name[i] > 32 {
			return true
		}
	}
	return false
}
