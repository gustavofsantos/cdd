package cmd

import (
	"testing"
)

func TestFuzzyScore(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		query    string
		minScore float64 // minimum expected score
	}{
		{
			name:     "exact match",
			text:     "authentication",
			query:    "authentication",
			minScore: 0.95,
		},
		{
			name:     "substring match",
			text:     "user authentication system",
			query:    "auth",
			minScore: 0.5,
		},
		{
			name:     "case insensitive match",
			text:     "User Authentication",
			query:    "auth",
			minScore: 0.5,
		},
		{
			name:     "partial word match",
			text:     "JWT token",
			query:    "token",
			minScore: 0.5,
		},
		{
			name:     "no match",
			text:     "redis cache",
			query:    "sql",
			minScore: 0.0,
		},
		{
			name:     "fuzzy match with gaps",
			text:     "context driven development",
			query:    "cdd",
			minScore: 0.3,
		},
		{
			name:     "short query in long text",
			text:     "The system shall validate user authentication before granting access",
			query:    "auth",
			minScore: 0.4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := FuzzyScore(tt.text, tt.query)
			if score < tt.minScore {
				t.Errorf("FuzzyScore(%q, %q) = %f, expected >= %f", tt.text, tt.query, score, tt.minScore)
			}
		})
	}
}

func TestFuzzyScoreNoMatch(t *testing.T) {
	// Test cases that should definitely not match
	tests := []struct {
		text  string
		query string
	}{
		{"redis database", "postgresql"},
		{"user login", "logout"},
		{"api endpoint", "database"},
	}

	for _, tt := range tests {
		score := FuzzyScore(tt.text, tt.query)
		if score > 0.1 { // Allow tiny floating point differences
			t.Errorf("FuzzyScore(%q, %q) = %f, expected near 0", tt.text, tt.query, score)
		}
	}
}

func TestFuzzyScoreRanking(t *testing.T) {
	// Test that scores rank matches correctly
	query := "auth"
	tests := []struct {
		text  string
		name  string
		score float64
	}{
		// Exact match should score highest
		{"auth", "exact match", 1.0},
		// Substring match should score high
		{"authentication", "substring match", 0.9},
		{"user authentication", "substring in phrase", 0.85},
		// Character-level fuzzy match lower
		{"access management", "chars in order", 0.3},
		// No match should be near zero
		{"password reset", "no match", 0.0},
	}

	var scores []float64
	for _, tt := range tests {
		score := FuzzyScore(tt.text, query)
		scores = append(scores, score)
	}

	// Verify ordering: scores should be mostly decreasing
	// (allowing for floating point and fuzzy logic variations)
	if scores[0] < scores[1] {
		t.Errorf("Exact match (%.2f) should score higher than substring (%.2f)", scores[0], scores[1])
	}
	if scores[1] < scores[2]-0.1 { // Some tolerance for fuzzy logic
		t.Errorf("Substring match (%.2f) should score >= substring in phrase (%.2f)", scores[1], scores[2])
	}
	if scores[4] > 0.1 {
		t.Errorf("No match should score near 0, got %.2f", scores[4])
	}
}

func TestFuzzyScoreEmptyQuery(t *testing.T) {
	score := FuzzyScore("some text", "")
	if score != 0.0 {
		t.Errorf("Empty query should score 0, got %f", score)
	}
}

func TestFuzzyScoreEmptyText(t *testing.T) {
	score := FuzzyScore("", "query")
	if score != 0.0 {
		t.Errorf("Empty text should score 0, got %f", score)
	}
}

func TestFuzzyScoreCaseSensitivity(t *testing.T) {
	// Should be case-insensitive
	tests := []struct {
		text  string
		query string
	}{
		{"Authentication", "auth"},
		{"AUTHENTICATION", "auth"},
		{"authentication", "AUTH"},
		{"AuThEnTiCaTiOn", "aUtHeNtIcAtIoN"},
	}

	for _, tt := range tests {
		score := FuzzyScore(tt.text, tt.query)
		if score < 0.65 {
			t.Errorf("Case variants should match: FuzzyScore(%q, %q) = %f", tt.text, tt.query, score)
		}
	}
}

func TestFuzzyScoreLongVsShortText(t *testing.T) {
	// Same query in different length texts
	query := "user"
	shortText := "user"
	longText := "The system manages users and user authentication across platforms"

	shortScore := FuzzyScore(shortText, query)
	longScore := FuzzyScore(longText, query)

	// Short exact match should score higher
	if shortScore <= longScore {
		t.Errorf("Exact match (%.2f) should score higher than substring in long text (%.2f)", shortScore, longScore)
	}

	// Both should be reasonably high though
	if longScore < 0.5 {
		t.Errorf("Substring in reasonable length text should score >= 0.5, got %.2f", longScore)
	}
}

func TestFuzzyScoreSpecialCharacters(t *testing.T) {
	// Should handle markdown and special chars
	tests := []struct {
		text  string
		query string
		name  string
	}{
		{"**bold text**", "bold", "markdown bold"},
		{"_italic_", "italic", "markdown italic"},
		{"`code`", "code", "markdown code"},
		{"[link](url)", "link", "markdown link"},
		{"The system `shall` validate.", "shall", "backtick code"},
	}

	for _, tt := range tests {
		score := FuzzyScore(tt.text, tt.query)
		if score < 0.5 {
			t.Errorf("%s: FuzzyScore(%q, %q) = %f, expected >= 0.5", tt.name, tt.text, tt.query, score)
		}
	}
}

func TestFuzzyScorePartialWords(t *testing.T) {
	// Partial word matches should work
	tests := []struct {
		text  string
		query string
	}{
		{"validation", "valid"},
		{"authenticate", "auth"},
		{"architectural", "architect"},
	}

	for _, tt := range tests {
		score := FuzzyScore(tt.text, tt.query)
		if score < 0.5 {
			t.Errorf("Partial word should match: FuzzyScore(%q, %q) = %f", tt.text, tt.query, score)
		}
	}
}
