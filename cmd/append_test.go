package stoic

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	s "github.com/skatkov/stoic/internal"
	"github.com/stretchr/testify/assert"
)

func TestAppendCommandAppendsText(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := s.NewContext(tmpDir, "", "", "")

	now := time.Date(2024, time.January, 5, 12, 0, 0, 0, time.UTC)
	cmd := NewAppendCommand(ctx, now, "quick note", false)
	cmd.Run()

	entryPath := filepath.Join(tmpDir, "2024-01-05.md")
	content, err := os.ReadFile(entryPath)
	assert.NoError(t, err)
	assert.Equal(t, "quick note\n", string(content))
}

func TestAppendCommandSupportsBullet(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := s.NewContext(tmpDir, "", "", "")

	now := time.Date(2024, time.January, 5, 9, 14, 0, 0, time.UTC)
	cmd := NewAppendCommand(ctx, now, "meeting notes", true)
	cmd.Run()

	entryPath := filepath.Join(tmpDir, "2024-01-05.md")
	content, err := os.ReadFile(entryPath)
	assert.NoError(t, err)
	assert.Equal(t, "- meeting notes\n", string(content))
}
