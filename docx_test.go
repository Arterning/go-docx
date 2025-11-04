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

func TestExtractTextWithOptions_FileNotFound(t *testing.T) {
	_, err := ExtractTextWithOptions("nonexistent.docx", Options{})
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

func TestOpenWithOptions_FileNotFound(t *testing.T) {
	_, err := OpenWithOptions("nonexistent.docx", Options{})
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

func TestGetHeadingLevel(t *testing.T) {
	tests := []struct {
		style    string
		expected int
	}{
		{"Heading1", 1},
		{"heading1", 1},
		{"Heading 1", 1},
		{"1", 1},
		{"Heading2", 2},
		{"heading2", 2},
		{"2", 2},
		{"Heading3", 3},
		{"Heading9", 9},
		{"Normal", 0},
		{"", 0},
		{"SomeOtherStyle", 0},
	}

	for _, tt := range tests {
		t.Run(tt.style, func(t *testing.T) {
			result := getHeadingLevel(tt.style)
			if result != tt.expected {
				t.Errorf("getHeadingLevel(%q) = %d, expected %d", tt.style, result, tt.expected)
			}
		})
	}
}
