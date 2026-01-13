package cmd

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var trackTemplates embed.FS

type trackData struct {
	TrackName string
	CreatedAt string
}

func renderTrackTemplate(name string, nameInFS string, data trackData) ([]byte, error) {
	tmpl, err := template.ParseFS(trackTemplates, "templates/"+nameInFS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %w", name, err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template %s: %w", name, err)
	}
	return buf.Bytes(), nil
}

func NewStartCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "start [track-name]",
		Short: "Create an isolated workspace (Track).",
		Long: `Creates an isolated workspace following the Lean CDD v4.1 protocol.
Usage: cdd start <track-name>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			trackDir := filepath.Join(".context/tracks", trackName)

			if _, err := fs.Stat(trackDir); !os.IsNotExist(err) {
				return fmt.Errorf("Error: Track '%s' exists.", trackName)
			}

			if err := fs.MkdirAll(trackDir, 0755); err != nil {
				return fmt.Errorf("Error creating track directory: %v", err)
			}

			data := trackData{
				TrackName: trackName,
				CreatedAt: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
			}

			// Spec Template
			specContent, err := renderTrackTemplate("spec.md", "spec.md", data)
			if err != nil {
				return err
			}
			if err := fs.WriteFile(filepath.Join(trackDir, "spec.md"), specContent, 0644); err != nil {
				return fmt.Errorf("failed to write spec.md: %w", err)
			}

			// Plan Template
			planContent, err := renderTrackTemplate("plan.md", "plan.md", data)
			if err != nil {
				return err
			}
			if err := fs.WriteFile(filepath.Join(trackDir, "plan.md"), planContent, 0644); err != nil {
				return fmt.Errorf("failed to write plan.md: %w", err)
			}

			// Decisions Log
			decisionsContent, err := renderTrackTemplate("decisions.md", "decisions.md", data)
			if err != nil {
				return err
			}
			if err := fs.WriteFile(filepath.Join(trackDir, "decisions.md"), decisionsContent, 0644); err != nil {
				return fmt.Errorf("failed to write decisions.md: %w", err)
			}

			// Metadata for Time Tracking (Internal)
			metadata := map[string]interface{}{
				"started_at": time.Now().Format(time.RFC3339),
			}
			metaBytes, err := json.MarshalIndent(metadata, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal metadata: %w", err)
			}
			if err := fs.WriteFile(filepath.Join(trackDir, "metadata.json"), metaBytes, 0644); err != nil {
				return fmt.Errorf("failed to write metadata.json: %w", err)
			}

			cmd.Printf("Track '%s' initialized.\n活跃 (Active) Track created with 3-file structure.\n", trackName)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewStartCmd(platform.NewRealFileSystem()))
}
