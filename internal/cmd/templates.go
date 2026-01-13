package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
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
