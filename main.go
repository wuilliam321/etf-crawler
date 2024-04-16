package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wuilliam321/etf-crawler/pkg/crawler"
)

func main() {
	// Check if a filename is provided as command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <format>")
		return
	}

	// Get the filename from command-line arguments
	format := os.Args[1]

	data, err := os.ReadFile("output.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(crawler.ToString(format, string(data)))
}
