# API Reference

## Functions

### ExtractText

```go
func ExtractText(filename string) (string, error)
```

Convenience function to extract text from a DOCX file with default options.

**Parameters:**
- `filename`: Path to the DOCX file

**Returns:**
- Extracted text content
- Error if file cannot be opened or parsed

**Example:**
```go
text, err := docx.ExtractText("document.docx")
```

---

### ExtractTextWithOptions

```go
func ExtractTextWithOptions(filename string, opts Options) (string, error)
```

Extracts text from a DOCX file with custom options.

**Parameters:**
- `filename`: Path to the DOCX file
- `opts`: Configuration options

**Returns:**
- Extracted text content
- Error if file cannot be opened or parsed

**Example:**
```go
text, err := docx.ExtractTextWithOptions("document.docx", docx.Options{
    ConvertHeadingsToMarkdown: true,
})
```

---

### Open

```go
func Open(filename string) (*Document, error)
```

Opens a DOCX file and returns a Document with default options.

**Parameters:**
- `filename`: Path to the DOCX file

**Returns:**
- Document object
- Error if file cannot be opened or parsed

**Example:**
```go
doc, err := docx.Open("document.docx")
if err != nil {
    log.Fatal(err)
}
defer doc.Close()
```

---

### OpenWithOptions

```go
func OpenWithOptions(filename string, opts Options) (*Document, error)
```

Opens a DOCX file with custom options.

**Parameters:**
- `filename`: Path to the DOCX file
- `opts`: Configuration options

**Returns:**
- Document object
- Error if file cannot be opened or parsed

**Example:**
```go
doc, err := docx.OpenWithOptions("document.docx", docx.Options{
    ConvertHeadingsToMarkdown: true,
})
if err != nil {
    log.Fatal(err)
}
defer doc.Close()
```

---

## Types

### Options

Configuration options for parsing DOCX documents.

```go
type Options struct {
    ConvertHeadingsToMarkdown bool
}
```

**Fields:**
- `ConvertHeadingsToMarkdown`: When `true`, converts Word heading styles (Heading 1-9) to Markdown format (e.g., `# Title`, `## Subtitle`). Default: `false`

---

### Document

Represents a DOCX document.

```go
type Document struct {
    // internal fields
}
```

**Methods:**

#### Text

```go
func (d *Document) Text() string
```

Returns the extracted text content.

**Example:**
```go
doc, _ := docx.Open("document.docx")
defer doc.Close()
text := doc.Text()
```

#### Close

```go
func (d *Document) Close() error
```

Closes the document and releases resources. Should be called when done with the document.

**Example:**
```go
doc, _ := docx.Open("document.docx")
defer doc.Close()
```

---

## Heading Style Mapping

When `ConvertHeadingsToMarkdown` is enabled, the following Word styles are recognized and converted:

| Word Style | Markdown Output |
|-----------|----------------|
| Heading 1, Heading1, 1 | `# Text` |
| Heading 2, Heading2, 2 | `## Text` |
| Heading 3, Heading3, 3 | `### Text` |
| Heading 4, Heading4, 4 | `#### Text` |
| Heading 5, Heading5, 5 | `##### Text` |
| Heading 6, Heading6, 6 | `###### Text` |
| Heading 7, Heading7, 7 | `####### Text` |
| Heading 8, Heading8, 8 | `######## Text` |
| Heading 9, Heading9, 9 | `######### Text` |
