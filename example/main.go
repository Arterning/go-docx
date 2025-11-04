package main

import (
	"fmt"
	"log"

	"github.com/Arterning/go-docx"
)

func main() {
	// Example 1: Simple usage - extract text directly
	fmt.Println("=== Example 1: Simple ExtractText ===")
	text, err := docx.ExtractText("sample.docx")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Println(text)
	}

	fmt.Println("\n=== Example 2: Using Document object ===")
	// Example 2: Using Document object
	doc, err := docx.Open("sample.docx")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer doc.Close()

	text = doc.Text()
	fmt.Println(text)
}
