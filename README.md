# go-docx

A simple and easy-to-use Go SDK for extracting text content from DOCX documents.

## Features

- Extract text content from DOCX files
- Preserve paragraph breaks with empty lines
- Extract table content (cells separated by commas, rows by newlines)
- Extract list content
- Simple and intuitive API

## Installation

```bash
go get github.com/Arterning/go-docx
```

## Usage

### Method 1: Simple extraction

```go
package main

import (
    "fmt"
    "log"

    "github.com/Arterning/go-docx"
)

func main() {
    text, err := docx.ExtractText("document.docx")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(text)
}
```

### Method 2: Using Document object

```go
package main

import (
    "fmt"
    "log"

    "github.com/Arterning/go-docx"
)

func main() {
    doc, err := docx.Open("document.docx")
    if err != nil {
        log.Fatal(err)
    }
    defer doc.Close()

    text := doc.Text()
    fmt.Println(text)
}
```

## Text Formatting

- **Paragraphs**: Separated by empty lines (double newline)
- **Tables**:
  - Cells in the same row are separated by commas
  - Rows are separated by newlines
- **Lists**: Extracted as regular text content

## Example Output

For a document containing:
- Two paragraphs
- A table with 2 rows and 3 columns

Output:
```
First paragraph text.

Second paragraph text.

Cell1,Cell2,Cell3
Cell4,Cell5,Cell6

```

## License

MIT
