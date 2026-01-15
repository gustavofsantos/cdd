package cmd

import (
	"encoding/json"
	"testing"
)

func TestPackToolboxDescribeIncludesLimit(t *testing.T) {
	// Test that the pack toolbox wrapper describes the limit parameter
	// We'll check this by examining the help output
	fs := NewPackCmd(nil)
	if fs == nil {
		t.Fatalf("NewPackCmd returned nil")
	}

	// Check that limit flag exists
	limitFlag := fs.Flags().Lookup("limit")
	if limitFlag == nil {
		t.Errorf("limit flag not found in pack command")
	}

	// Check flag short name
	if limitFlag.Shorthand != "l" {
		t.Errorf("limit flag shorthand should be 'l', got %q", limitFlag.Shorthand)
	}
}

func TestPackToolboxLimitParameter(t *testing.T) {
	// Test that limit can be passed through command
	fs := NewPackCmd(nil)

	// Try setting limit flag
	err := fs.Flags().Set("limit", "10")
	if err != nil {
		t.Errorf("Failed to set limit flag: %v", err)
	}

	// Verify it was set
	limitVal, err := fs.Flags().GetInt("limit")
	if err != nil {
		t.Errorf("Failed to get limit flag: %v", err)
	}

	if limitVal != 10 {
		t.Errorf("Expected limit 10, got %d", limitVal)
	}
}

func TestPackToolboxLimitWithFocus(t *testing.T) {
	// Test that limit flag works with focus flag
	fs := NewPackCmd(nil)

	err := fs.Flags().Set("focus", "log")
	if err != nil {
		t.Errorf("Failed to set focus flag: %v", err)
	}

	err = fs.Flags().Set("limit", "5")
	if err != nil {
		t.Errorf("Failed to set limit flag: %v", err)
	}

	focusVal, _ := fs.Flags().GetString("focus")
	limitVal, _ := fs.Flags().GetInt("limit")

	if focusVal != "log" {
		t.Errorf("Expected focus 'log', got %q", focusVal)
	}

	if limitVal != 5 {
		t.Errorf("Expected limit 5, got %d", limitVal)
	}
}

func TestPackToolboxLimitJSON(t *testing.T) {
	// Test that limit parameter can be serialized/deserialized as JSON
	// This simulates what Amp's toolbox would do
	params := map[string]interface{}{
		"focus": "command",
		"limit": 3,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		t.Errorf("Failed to marshal params: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal params: %v", err)
	}

	if unmarshaled["focus"] != "command" {
		t.Errorf("focus not preserved in JSON")
	}

	if unmarshaled["limit"] != float64(3) {
		t.Errorf("limit not preserved in JSON, got %v", unmarshaled["limit"])
	}
}

func TestPackToolboxLimitZeroJSON(t *testing.T) {
	// Test that limit=0 can be serialized
	params := map[string]interface{}{
		"focus": "log",
		"limit": 0,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		t.Errorf("Failed to marshal params with limit=0: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal params: %v", err)
	}

	if unmarshaled["limit"] != float64(0) {
		t.Errorf("limit=0 not preserved, got %v", unmarshaled["limit"])
	}
}

func TestPackToolboxNoLimit(t *testing.T) {
	// Test that no limit parameter defaults to -1
	fs := NewPackCmd(nil)

	limitVal, err := fs.Flags().GetInt("limit")
	if err != nil {
		t.Errorf("Failed to get limit flag: %v", err)
	}

	if limitVal != -1 {
		t.Errorf("Default limit should be -1, got %d", limitVal)
	}
}
