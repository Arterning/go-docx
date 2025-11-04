package docx

import (
	"testing"
)

func TestExtractText_FileNotFound(t *testing.T) {
	_, err := ExtractText("nonexistent.docx")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestOpen_FileNotFound(t *testing.T) {
	_, err := Open("nonexistent.docx")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestDocument_Close(t *testing.T) {
	doc := &Document{}
	err := doc.Close()
	if err != nil {
		t.Errorf("Close() should not return error for nil zipReader, got: %v", err)
	}
}

func TestDocument_Text(t *testing.T) {
	doc := &Document{
		content: "test content",
	}

	text := doc.Text()
	if text != "test content" {
		t.Errorf("Expected 'test content', got '%s'", text)
	}
}
