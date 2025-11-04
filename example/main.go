package main

import (
	"fmt"
	"log"

	"github.com/Arterning/go-docx"
)

func main() {
	// Example 1: Simple usage - extract text directly
	fmt.Println("=== Example 1: Simple ExtractText ===")
	text, err := docx.ExtractTextWithOptions("sample.docx", docx.Options{
        ConvertHeadingsToMarkdown: true,
    })
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

	fmt.Println("\n=== Example 3: Convert headings to Markdown ===")
	// Example 3: With options to convert headings to Markdown
	text, err = docx.ExtractTextWithOptions("sample.docx", docx.Options{
		ConvertHeadingsToMarkdown: true,
	})
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Println(text)
	}

	fmt.Println("\n=== Example 4: Using OpenWithOptions ===")
	// Example 4: Using Document object with options
	doc2, err := docx.OpenWithOptions("sample.docx", docx.Options{
		ConvertHeadingsToMarkdown: true,
	})
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer doc2.Close()

	text = doc2.Text()
	fmt.Println(text)
}
