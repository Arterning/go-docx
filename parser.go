package docx

import (
	"encoding/xml"
	"strings"
)

// XML structure definitions for DOCX
type wDocument struct {
	XMLName xml.Name `xml:"document"`
	Body    wBody    `xml:"body"`
}

type wBody struct {
	Items []interface{} `xml:",any"`
}

type wParagraph struct {
	XMLName xml.Name `xml:"p"`
	Runs    []wRun   `xml:"r"`
}

type wRun struct {
	XMLName xml.Name `xml:"r"`
	Text    []wText  `xml:"t"`
}

type wText struct {
	XMLName xml.Name `xml:"t"`
	Text    string   `xml:",chardata"`
}

type wTable struct {
	XMLName xml.Name `xml:"tbl"`
	Rows    []wRow   `xml:"tr"`
}

type wRow struct {
	XMLName xml.Name `xml:"tr"`
	Cells   []wCell  `xml:"tc"`
}

type wCell struct {
	XMLName    xml.Name     `xml:"tc"`
	Paragraphs []wParagraph `xml:"p"`
}

// parseDocumentXML parses the document.xml content and extracts text
func parseDocumentXML(data []byte, opts Options) (string, error) {
	// Decode raw XML to process elements
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	var result strings.Builder
	var inParagraph bool
	var inTable bool
	var inParagraphProperties bool
	var currentTableRow []string
	var tableRows [][]string
	var paragraphText strings.Builder
	var paragraphStyle string

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "p":
				inParagraph = true
				paragraphText.Reset()
				paragraphStyle = ""
			case "pPr":
				// Paragraph properties
				inParagraphProperties = true
			case "pStyle":
				// Paragraph style - extract the value
				if inParagraphProperties {
					for _, attr := range elem.Attr {
						if attr.Name.Local == "val" {
							paragraphStyle = attr.Value
							break
						}
					}
				}
			case "tbl":
				inTable = true
				tableRows = nil
			case "tr":
				if inTable {
					currentTableRow = nil
				}
			case "tc":
				// Table cell - will collect text from paragraphs inside
			case "t":
				// Text element - read the content
				var text wText
				if err := decoder.DecodeElement(&text, &elem); err == nil {
					if inParagraph {
						paragraphText.WriteString(text.Text)
					}
				}
			}

		case xml.EndElement:
			switch elem.Name.Local {
			case "pPr":
				inParagraphProperties = false

			case "p":
				if inParagraph && !inTable {
					// Regular paragraph - add text with double newline
					text := strings.TrimSpace(paragraphText.String())
					if text != "" {
						// Check if this is a heading and convert to markdown if needed
						headingLevel := getHeadingLevel(paragraphStyle)
						if headingLevel > 0 && opts.ConvertHeadingsToMarkdown {
							// Add markdown heading prefix
							result.WriteString(strings.Repeat("#", headingLevel))
							result.WriteString(" ")
							result.WriteString(text)
							result.WriteString("\n\n")
						} else {
							result.WriteString(text)
							result.WriteString("\n\n")
						}
					}
				} else if inTable {
					// Paragraph inside table cell
					text := strings.TrimSpace(paragraphText.String())
					if text != "" {
						currentTableRow = append(currentTableRow, text)
					}
				}
				inParagraph = false
				paragraphText.Reset()
				paragraphStyle = ""

			case "tr":
				if inTable && len(currentTableRow) > 0 {
					tableRows = append(tableRows, currentTableRow)
					currentTableRow = nil
				}

			case "tc":
				// End of table cell - paragraph text already added to currentTableRow
				paragraphText.Reset()

			case "tbl":
				// End of table - format and add all rows
				if len(tableRows) > 0 {
					for _, row := range tableRows {
						result.WriteString(strings.Join(row, ","))
						result.WriteString("\n")
					}
					result.WriteString("\n")
				}
				inTable = false
				tableRows = nil
			}
		}
	}

	// Remove trailing newlines
	return strings.TrimRight(result.String(), "\n"), nil
}

// getHeadingLevel returns the heading level (1-9) from a paragraph style, or 0 if not a heading
func getHeadingLevel(style string) int {
	if style == "" {
		return 0
	}

	// Handle both "Heading1" and "1" style formats
	style = strings.ToLower(style)

	// Common Word heading styles
	headingMap := map[string]int{
		"heading1": 1, "heading 1": 1, "1": 1,
		"heading2": 2, "heading 2": 2, "2": 2,
		"heading3": 3, "heading 3": 3, "3": 3,
		"heading4": 4, "heading 4": 4, "4": 4,
		"heading5": 5, "heading 5": 5, "5": 5,
		"heading6": 6, "heading 6": 6, "6": 6,
		"heading7": 7, "heading 7": 7, "7": 7,
		"heading8": 8, "heading 8": 8, "8": 8,
		"heading9": 9, "heading 9": 9, "9": 9,
	}

	if level, ok := headingMap[style]; ok {
		return level
	}

	return 0
}
