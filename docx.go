package docx

import (
	"archive/zip"
	"fmt"
	"io"
)

// Document represents a DOCX document
type Document struct {
	zipReader *zip.ReadCloser
	content   string
}

// Open opens a DOCX file and returns a Document
func Open(filename string) (*Document, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open docx file: %w", err)
	}

	doc := &Document{
		zipReader: r,
	}

	if err := doc.parse(); err != nil {
		r.Close()
		return nil, err
	}

	return doc, nil
}

// ExtractText is a convenience function to extract text from a DOCX file
func ExtractText(filename string) (string, error) {
	doc, err := Open(filename)
	if err != nil {
		return "", err
	}
	defer doc.Close()

	return doc.Text(), nil
}

// Text returns the extracted text content
func (d *Document) Text() string {
	return d.content
}

// Close closes the document
func (d *Document) Close() error {
	if d.zipReader != nil {
		return d.zipReader.Close()
	}
	return nil
}

// parse extracts and parses the document.xml file
func (d *Document) parse() error {
	// Find and read word/document.xml
	var documentXML *zip.File
	for _, f := range d.zipReader.File {
		if f.Name == "word/document.xml" {
			documentXML = f
			break
		}
	}

	if documentXML == nil {
		return fmt.Errorf("word/document.xml not found in docx file")
	}

	rc, err := documentXML.Open()
	if err != nil {
		return fmt.Errorf("failed to open document.xml: %w", err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("failed to read document.xml: %w", err)
	}

	// Parse the XML and extract text
	text, err := parseDocumentXML(data)
	if err != nil {
		return fmt.Errorf("failed to parse document.xml: %w", err)
	}

	d.content = text
	return nil
}
