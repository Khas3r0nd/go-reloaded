package main

import (
	"fmt"
	"go-reloaded/strproc"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("Incorrect amount of arguments %v, expected 2\nUsage: go run main.go inputFile.txt outputFile.txt\n", len(args))
		os.Exit(1)
	}
	err := strproc.Process(args[0], args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
