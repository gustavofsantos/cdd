package cmd

import (
	"testing"
)

func TestExtractParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single paragraph",
			input:    `This is a paragraph.`,
			expected: []string{"This is a paragraph."},
		},
		{
			name: "multiple paragraphs separated by blank lines",
			input: `First paragraph.

Second paragraph.

Third paragraph.`,
			expected: []string{
				"First paragraph.",
				"Second paragraph.",
				"Third paragraph.",
			},
		},
		{
			name: "multiline paragraph",
			input: `Line one
Line two
Line three

Second paragraph.`,
			expected: []string{
				"Line one\nLine two\nLine three",
				"Second paragraph.",
			},
		},
		{
			name: "trailing and leading whitespace",
			input: `  
First paragraph.

Second paragraph.
  `,
			expected: []string{
				"First paragraph.",
				"Second paragraph.",
			},
		},
		{
			name:     "empty input",
			input:    ``,
			expected: []string{},
		},
		{
			name: "only whitespace",
			input: `   

   
`,
			expected: []string{},
		},
		{
			name: "markdown header",
			input: `# Title

Some content under the title.`,
			expected: []string{
				"# Title",
				"Some content under the title.",
			},
		},
		{
			name: "multiple headers with content",
			input: `# Main Header

Content for main section.

## Subsection

Content for subsection.`,
			expected: []string{
				"# Main Header",
				"Content for main section.",
				"## Subsection",
				"Content for subsection.",
			},
		},
		{
			name: "markdown list items are one paragraph",
			input: `- Item one
- Item two
- Item three

Next paragraph.`,
			expected: []string{
				"- Item one\n- Item two\n- Item three",
				"Next paragraph.",
			},
		},
		{
			name:  "code block with triple backticks",
			input: "Here is some code:\n\n```go\nfunc main() {\n    fmt.Println(\"hello\")\n}\n```\n\nMore text after code.",
			expected: []string{
				"Here is some code:",
				"```go\nfunc main() {\n    fmt.Println(\"hello\")\n}\n```",
				"More text after code.",
			},
		},
		{
			name:  "inline code in paragraph",
			input: "The `extract` function works well.\n\nNext paragraph.",
			expected: []string{
				"The `extract` function works well.",
				"Next paragraph.",
			},
		},
		{
			name: "bold and italic text",
			input: `This is **bold** and *italic* text.

Another paragraph with ***bold italic***.`,
			expected: []string{
				"This is **bold** and *italic* text.",
				"Another paragraph with ***bold italic***.",
			},
		},
		{
			name: "nested requirements section",
			input: `## Requirements

- Ubiquitous: The system shall work.
- Event-driven: When triggered, respond.
- State-driven: While active, persist state.

## Next Section

Content here.`,
			expected: []string{
				"## Requirements",
				"- Ubiquitous: The system shall work.\n- Event-driven: When triggered, respond.\n- State-driven: While active, persist state.",
				"## Next Section",
				"Content here.",
			},
		},
		{
			name: "blockquote",
			input: `Regular paragraph.

> This is a blockquote
> spanning multiple lines.

After blockquote.`,
			expected: []string{
				"Regular paragraph.",
				"> This is a blockquote\n> spanning multiple lines.",
				"After blockquote.",
			},
		},
		{
			name: "horizontal rule",
			input: `Text before.

---

Text after.`,
			expected: []string{
				"Text before.",
				"---",
				"Text after.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractParagraphs(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ExtractParagraphs() returned %d paragraphs, expected %d", len(result), len(tt.expected))
			}
			for i, para := range result {
				if para != tt.expected[i] {
					t.Errorf("Paragraph %d: got %q, expected %q", i, para, tt.expected[i])
				}
			}
		})
	}
}
