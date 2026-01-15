package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"cdd/internal/platform"
)

// ExtractParagraphs splits markdown text by blank lines.
// A paragraph is one or more consecutive non-empty lines.
// Trailing/leading whitespace is stripped from the entire input and each paragraph.
func ExtractParagraphs(text string) []string {
	lines := strings.Split(text, "\n")
	
	var paragraphs []string
	var currentParagraph []string
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		
		if trimmed == "" {
			// Blank line: end current paragraph if it exists
			if len(currentParagraph) > 0 {
				paragraphs = append(paragraphs, strings.Join(currentParagraph, "\n"))
				currentParagraph = []string{}
			}
		} else {
			// Non-empty line: add to current paragraph
			currentParagraph = append(currentParagraph, line)
		}
	}
	
	// Don't forget the last paragraph
	if len(currentParagraph) > 0 {
		paragraphs = append(paragraphs, strings.Join(currentParagraph, "\n"))
	}
	
	return paragraphs
}

// FuzzyScore returns a similarity score between text and query (0.0 to 1.0).
// Uses a combination of substring matching and character-level matching.
// Higher scores indicate better matches.
func FuzzyScore(text, query string) float64 {
	if query == "" {
		return 0.0
	}

	textLower := strings.ToLower(text)
	queryLower := strings.ToLower(query)

	// Exact match
	if textLower == queryLower {
		return 1.0
	}

	// Substring match
	if strings.Contains(textLower, queryLower) {
		// Higher score for substring matches: query is a contiguous sequence
		idx := strings.Index(textLower, queryLower)
		// Base score is high for substring matches
		baseScore := 0.65
		// Bonus based on how early the match appears (earlier is better)
		positionBonus := (1.0 - (float64(idx) / float64(len(textLower)))) * 0.25
		return baseScore + positionBonus
	}

	// Character-level fuzzy matching
	score := fuzzyCharMatch(textLower, queryLower)
	return score
}

// fuzzyCharMatch performs character-by-character matching with gaps allowed.
// Returns a score based on how many query characters appear in order in the text.
func fuzzyCharMatch(text, query string) float64 {
	if len(query) == 0 {
		return 0.0
	}
	if len(text) == 0 {
		return 0.0
	}

	// Count how many characters from query appear in text in order
	queryIdx := 0
	textIdx := 0
	matchedChars := 0

	for textIdx < len(text) && queryIdx < len(query) {
		if text[textIdx] == query[queryIdx] {
			matchedChars++
			queryIdx++
		}
		textIdx++
	}

	if matchedChars == 0 {
		return 0.0
	}

	// Score based on matched characters and how compact they are
	charMatchRatio := float64(matchedChars) / float64(len(query))
	if matchedChars != len(query) {
		// Penalize for gaps - matches must be complete
		return 0.0
	}

	// Bonus for how early the match completes
	positionPenalty := float64(textIdx) / float64(len(text))
	return charMatchRatio * (1.0 - positionPenalty*0.3)
}

// Spec represents a specification file with its metadata and content.
type Spec struct {
	Name    string // filename (e.g., "log.md")
	Path    string // full path to the file
	Content []byte // file contents
}

// DiscoverSpecs reads all markdown files from a directory and returns them as Spec objects.
func DiscoverSpecs(fs platform.FileSystem, dir string) ([]Spec, error) {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading specs directory: %v", err)
	}

	var specs []Spec

	for _, entry := range entries {
		// Skip directories, only process files
		if entry.IsDir() {
			continue
		}

		// Only include .md files
		if !platform.EndsWithString(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		content, err := fs.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading spec file %s: %v", filePath, err)
		}

		specs = append(specs, Spec{
			Name:    entry.Name(),
			Path:    filePath,
			Content: content,
		})
	}

	return specs, nil
}

// ParagraphMatch represents a paragraph that matched a search query.
type ParagraphMatch struct {
	Paragraph string  // the matching paragraph text
	SpecName  string  // the spec file it came from
	Score     float64 // the fuzzy match score (0.0-1.0)
}

// FilterParagraphs searches all specs for paragraphs matching the given topic.
// Returns matches sorted by score (highest first).
func FilterParagraphs(specs []Spec, topic string, minScore float64) []ParagraphMatch {
	if topic == "" || len(specs) == 0 {
		return []ParagraphMatch{}
	}

	var matches []ParagraphMatch

	for _, spec := range specs {
		// Extract paragraphs from spec content
		content := string(spec.Content)
		paragraphs := ExtractParagraphs(content)

		// Score each paragraph against the topic
		for _, para := range paragraphs {
			score := FuzzyScore(para, topic)
			if score >= minScore {
				matches = append(matches, ParagraphMatch{
					Paragraph: para,
					SpecName:  spec.Name,
					Score:     score,
				})
			}
		}
	}

	// Sort matches by score (highest first)
	sortMatches(matches)

	return matches
}

// sortMatches sorts ParagraphMatch slice by score in descending order
func sortMatches(matches []ParagraphMatch) {
	// Simple bubble sort for small lists
	for i := 0; i < len(matches); i++ {
		for j := i + 1; j < len(matches); j++ {
			if matches[j].Score > matches[i].Score {
				matches[i], matches[j] = matches[j], matches[i]
			}
		}
	}
}
