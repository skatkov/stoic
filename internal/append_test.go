package stoic

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppendCreatesFileWithTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	templatePath := filepath.Join(tmpDir, "template.md")
	err := os.WriteFile(templatePath, []byte("heading\n"), 0o644)
	assert.NoError(t, err)

	ctx := NewContext(tmpDir, "", "", templatePath)
	entry := NewEntry(ctx, time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC))

	err = ctx.Append(entry, "note content")
	assert.NoError(t, err)

	content, err := os.ReadFile(entry.Filepath())
	assert.NoError(t, err)
	assert.Equal(t, "heading\n\nnote content\n", string(content))
}

func TestAppendAddsNewlineWhenMissing(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := NewContext(tmpDir, "", "", "")
	entry := NewEntry(ctx, time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC))

	err := os.WriteFile(entry.Filepath(), []byte("first line"), 0o644)
	assert.NoError(t, err)

	err = ctx.Append(entry, "second line")
	assert.NoError(t, err)

	data, err := os.ReadFile(entry.Filepath())
	assert.NoError(t, err)
	assert.Equal(t, "first line\n\nsecond line\n", string(data))
}

func TestAppendOnEmptyFileDoesNotPrependNewline(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := NewContext(tmpDir, "", "", "")
	entry := NewEntry(ctx, time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC))

	err := ctx.Append(entry, "only line")
	assert.NoError(t, err)

	data, err := os.ReadFile(entry.Filepath())
	assert.NoError(t, err)
	assert.Equal(t, "only line\n", string(data))
}
