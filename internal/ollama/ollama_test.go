package ollama

import (
	"testing"
)

func TestParseResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "basic response",
			input: `feat(auth): add login functionality
fix(ui): resolve button alignment
refactor(db): optimize query performance`,
			expected: []string{
				"feat(auth): add login functionality",
				"fix(ui): resolve button alignment",
				"refactor(db): optimize query performance",
			},
		},
		{
			name: "response with numbers",
			input: `1. feat(auth): add login
2. fix(ui): fix button
3. refactor(db): optimize`,
			expected: []string{
				"feat(auth): add login",
				"fix(ui): fix button",
				"refactor(db): optimize",
			},
		},
		{
			name: "response with markdown",
			input: `- feat(auth): add login
- fix(ui): fix button
- refactor(db): optimize`,
			expected: []string{
				"feat(auth): add login",
				"fix(ui): fix button",
				"refactor(db): optimize",
			},
		},
		{
			name: "empty lines",
			input: `feat(auth): add login

fix(ui): fix button

refactor(db): optimize`,
			expected: []string{
				"feat(auth): add login",
				"fix(ui): fix button",
				"refactor(db): optimize",
			},
		},
		{
			name:     "empty response",
			input:    "",
			expected: []string{},
		},
		{
			name: "trailing period filtered",
			input: `feat(auth): add login.
fix(ui): fix button.`,
			expected: []string{},
		},
		{
			name: "too long filtered",
			input: `feat(auth): this is a very long commit message that exceeds the seventy two character limit and should be filtered out`,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseResponse(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("parseResponse() returned %d messages, want %d", len(result), len(tt.expected))
				for i, msg := range result {
					t.Logf("  got[%d]: %q", i, msg)
				}
				return
			}
			for i, msg := range result {
				if msg != tt.expected[i] {
					t.Errorf("parseResponse()[%d] = %q, want %q", i, msg, tt.expected[i])
				}
			}
		})
	}
}
