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

### Method 3: Convert headings to Markdown

```go
package main

import (
    "fmt"
    "log"

    "github.com/Arterning/go-docx"
)

func main() {
    // Extract text with headings converted to Markdown format
    text, err := docx.ExtractTextWithOptions("document.docx", docx.Options{
        ConvertHeadingsToMarkdown: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(text)
}
```

### Method 4: Using Document object with options

```go
package main

import (
    "fmt"
    "log"

    "github.com/Arterning/go-docx"
)

func main() {
    doc, err := docx.OpenWithOptions("document.docx", docx.Options{
        ConvertHeadingsToMarkdown: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer doc.Close()

    text := doc.Text()
    fmt.Println(text)
}
```

## Configuration Options

### ConvertHeadingsToMarkdown

When set to `true`, Word heading styles (Heading 1-9) are converted to Markdown format:
- Heading 1 → `# Title`
- Heading 2 → `## Title`
- Heading 3 → `### Title`
- etc.

When set to `false` (default), headings are treated as regular paragraphs.

## Text Formatting

- **Paragraphs**: Separated by empty lines (double newline)
- **Tables**:
  - Cells in the same row are separated by commas
  - Rows are separated by newlines
- **Lists**: Extracted as regular text content
- **Headings**: Can be converted to Markdown format (see Options above)

## Example Output

### Without Markdown conversion (default)

For a document containing:
- A Heading 1: "Introduction"
- Two paragraphs
- A table with 2 rows and 3 columns

Output:
```
Introduction

First paragraph text.

Second paragraph text.

Cell1,Cell2,Cell3
Cell4,Cell5,Cell6

```

### With Markdown conversion

Same document with `ConvertHeadingsToMarkdown: true`:

Output:
```
# Introduction

First paragraph text.

Second paragraph text.

Cell1,Cell2,Cell3
Cell4,Cell5,Cell6

```

## License

MIT
