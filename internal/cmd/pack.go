package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cdd/internal/platform"

	"github.com/charmbracelet/glamour"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var (
	packFocus string
	packRaw   bool
	packLimit int
)

func NewPackCmd(fs platform.FileSystem) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pack --focus <topic>",
		Short: "Pack relevant specs by topic focus.",
		Long: `Pack compresses global specs by extracting only paragraphs relevant to a given topic.
This helps manage cognitive load in large projects by delivering focused context.

The command searches all .context/specs/ files using fuzzy matching and outputs
only the paragraphs that match your topic, ranked by relevance.

FLAGS:
  -f, --focus <topic>   Topic to search for (required). E.g., 'log', 'command', 'auth'.
  -r, --raw            Output plain text without markdown rendering.
  -l, --limit <number>  Maximum paragraphs to return (default: no limit, use -1 for no limit).

EXAMPLES:
  $ cdd pack --focus log             # Find log-related content
  $ cdd pack --focus command         # Find command documentation
  $ cdd pack --focus auth --raw      # Raw output on authentication topics
  $ cdd pack --focus log --limit 5   # Limit results to 5 paragraphs`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPackCmd(cmd, fs)
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return GetPackCompletion(args, toComplete)
		},
	}

	cmd.Flags().StringVarP(&packFocus, "focus", "f", "", "Topic to search for (required)")
	cmd.Flags().BoolVarP(&packRaw, "raw", "r", false, "Output plain text without markdown rendering")
	cmd.Flags().IntVarP(&packLimit, "limit", "l", -1, "Maximum paragraphs to return (default: no limit)")

	return cmd
}

func runPackCmd(cmd *cobra.Command, fs platform.FileSystem) error {
	// Validate required flag
	if packFocus == "" {
		return fmt.Errorf("--focus flag is required. Example: cdd pack --focus log")
	}

	// Discover specs from .context/specs
	// Use absolute path or relative from project root
	specsPath := ".context/specs"
	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		// Try alternate path for tests
		specsPath = filepath.Join("..", "..", ".context", "specs")
		specs, err = DiscoverSpecs(fs, specsPath)
		if err != nil {
			return fmt.Errorf("error discovering specs: %v", err)
		}
	}

	if len(specs) == 0 {
		return fmt.Errorf("no spec files found in %s", specsPath)
	}

	// Filter paragraphs by topic with a reasonable threshold
	minScore := 0.5
	allMatches := FilterParagraphs(specs, packFocus, minScore)

	if len(allMatches) == 0 {
		fmt.Fprintf(cmd.OutOrStdout(), "No matches found for topic: %q\n", packFocus)
		fmt.Fprintf(cmd.OutOrStdout(), "Try other topics or check available specs with: cdd view\n")
		return nil
	}

	// Apply limit if specified
	matches, err := LimitResults(allMatches, packLimit)
	if err != nil {
		return fmt.Errorf("error applying limit: %v", err)
	}

	// Build markdown output (pass total count for truncation message)
	markdown := buildPackMarkdownWithLimit(matches, packFocus, len(allMatches))

	// Output: raw or rendered
	isTerminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	if packRaw || !isTerminal {
		// Plain text output
		fmt.Fprint(cmd.OutOrStdout(), markdown)
		return nil
	}

	// Rendered markdown output
	rendered, err := glamour.Render(markdown, "dark")
	if err != nil {
		// Fall back to raw if rendering fails
		fmt.Fprint(cmd.OutOrStdout(), markdown)
		return nil
	}

	fmt.Fprint(cmd.OutOrStdout(), rendered)
	return nil
}

// buildPackMarkdown constructs markdown output from filtered matches
func buildPackMarkdown(matches []ParagraphMatch, topic string) string {
	return buildPackMarkdownWithLimit(matches, topic, len(matches))
}

// buildPackMarkdownWithLimit constructs markdown output with limit information
func buildPackMarkdownWithLimit(matches []ParagraphMatch, topic string, totalCount int) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Context Pack: %q\n\n", topic))
	
	// Show if results were truncated
	if len(matches) == 0 && totalCount > 0 {
		sb.WriteString(fmt.Sprintf("Found %d relevant paragraphs across specs (showing 0 matches).\n\n", totalCount))
	} else if len(matches) < totalCount {
		sb.WriteString(fmt.Sprintf("Found %d relevant paragraphs across specs (showing %d of %d).\n\n", totalCount, len(matches), totalCount))
	} else {
		sb.WriteString(fmt.Sprintf("Found %d relevant paragraphs across specs.\n\n", len(matches)))
	}

	// Group by spec file
	specGroups := make(map[string][]ParagraphMatch)
	for _, match := range matches {
		specGroups[match.SpecName] = append(specGroups[match.SpecName], match)
	}

	// Output grouped by spec
	for _, spec := range getSortedSpecNames(specGroups) {
		matches := specGroups[spec]
		sb.WriteString(fmt.Sprintf("## From %s\n\n", spec))

		for i, match := range matches {
			sb.WriteString(fmt.Sprintf("**[Match %d] (Score: %.2f)**\n\n", i+1, match.Score))
			sb.WriteString(match.Paragraph)
			sb.WriteString("\n\n---\n\n")
		}
	}

	return sb.String()
}

// getSortedSpecNames returns spec file names in a consistent order
func getSortedSpecNames(groups map[string][]ParagraphMatch) []string {
	names := make([]string, 0, len(groups))
	for name := range groups {
		names = append(names, name)
	}
	// Simple bubble sort for consistent ordering
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			if names[j] < names[i] {
				names[i], names[j] = names[j], names[i]
			}
		}
	}
	return names
}

// GetPackCompletionTopics returns a list of suggested topics for pack command completion
func GetPackCompletionTopics() []string {
	return []string{
		"log",
		"view",
		"command",
		"specification",
		"requirement",
		"authentication",
		"authorization",
		"tracking",
		"decision",
		"architecture",
		"testing",
		"deployment",
		"configuration",
		"error",
		"validation",
	}
}

// GetPackCompletion provides shell completion suggestions for the pack command
func GetPackCompletion(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	topics := GetPackCompletionTopics()

	// Filter topics that start with toComplete (case-insensitive)
	toLower := strings.ToLower(toComplete)
	var filtered []string
	for _, topic := range topics {
		if strings.HasPrefix(strings.ToLower(topic), toLower) {
			filtered = append(filtered, topic)
		}
	}

	// Always return a slice, not nil
	if filtered == nil {
		filtered = []string{}
	}

	return filtered, cobra.ShellCompDirectiveNoFileComp
}

// LimitResults returns the first N items from a slice of ParagraphMatch.
// If limit is negative, returns all matches.
// If limit is zero, returns an empty slice.
// If limit is greater than available items, returns all items.
func LimitResults(matches []ParagraphMatch, limit int) ([]ParagraphMatch, error) {
	// If limit is negative, return all matches
	if limit < 0 {
		return matches, nil
	}

	// If limit is 0, return empty slice
	if limit == 0 {
		return []ParagraphMatch{}, nil
	}

	// If limit is greater than available, return all
	if limit >= len(matches) {
		return matches, nil
	}

	// Return first N matches
	return matches[:limit], nil
}

func init() {
	rootCmd.AddCommand(NewPackCmd(platform.NewRealFileSystem()))
}
