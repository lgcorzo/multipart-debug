package minio

import (
	"bytes"
	"os"
	"testing"
)

func TestHandleOutput(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "string input",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "struct input",
			input:    struct{ Status string }{"ok"},
			expected: "{\n  \"Status\": \"ok\"\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}
			os.Stdout = w

			handleOutput(tt.input)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			if _, err := buf.ReadFrom(r); err != nil {
				t.Fatalf("failed to read from pipe: %v", err)
			}

			if buf.String() != tt.expected {
				t.Errorf("got %q, want %q", buf.String(), tt.expected)
			}
		})
	}
}
