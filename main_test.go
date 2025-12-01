package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	cmd "github.com/skatkov/stoic/cmd"
	s "github.com/skatkov/stoic/internal"
	"github.com/stretchr/testify/assert"
)

func TestResolveAppendInputPrefersArgs(t *testing.T) {
	pr, pw, err := os.Pipe()
	assert.NoError(t, err)
	_, _ = pw.WriteString("from pipe")
	_ = pw.Close()
	defer pr.Close()

	text, ok, err := resolveAppendInput([]string{"from", "args"}, pr)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "from args", text)
}

func TestResolveAppendInputReadsPipe(t *testing.T) {
	pr, pw, err := os.Pipe()
	assert.NoError(t, err)
	_, _ = pw.WriteString("piped note")
	_ = pw.Close()
	defer pr.Close()

	text, ok, err := resolveAppendInput(nil, pr)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "piped note", strings.TrimSpace(text))
}

func TestResolveAppendInputNoInput(t *testing.T) {
	text, ok, err := resolveAppendInput(nil, nil)
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Empty(t, text)
}

func TestPipedAppendWritesEntry(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := s.NewContext(tmpDir, "", "", "")
	date := time.Date(2024, time.January, 5, 8, 0, 0, 0, time.UTC)

	pr, pw, err := os.Pipe()
	assert.NoError(t, err)
	_, _ = pw.WriteString("piped append")
	_ = pw.Close()
	defer pr.Close()

	text, ok, err := resolveAppendInput(nil, pr)
	assert.NoError(t, err)
	assert.True(t, ok)

	cmd.NewAppendCommand(ctx, date, text, false).Run()

	entryPath := filepath.Join(tmpDir, "2024-01-05.md")
	content, err := os.ReadFile(entryPath)
	assert.NoError(t, err)
	assert.Equal(t, "piped append\n", string(content))
}
