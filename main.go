package main

import (
	"fmt"
	"os"
)

func main() {
	var filename, query string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		} else if os.Args[i] == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		}
	}

	if filename == "" || query == "" {
		fmt.Println("Использование: ./lr3_go --file <путь_к_файлу> --query '<команда>'")
		os.Exit(1)
	}

	db := NewDatabase()
	db.LoadFromFile(filename)
	db.ProcessCommand(query)
	_ = db.SaveToFile(filename)
}
