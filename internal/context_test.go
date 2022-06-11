package stoic

import (
	"os"
	"os/user"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {
	current_dir, _ := os.Getwd()
	test_folder := strings.TrimSuffix(current_dir, "/internal") + "/test/data"

	ctx := NewContext(test_folder+"/journal_with_entry", "", "", "")
	assert.Equal(t, []string{test_folder + "/journal_with_entry/" + "2022-jan-01.txt"}, ctx.Files())

	ctx = NewContext(test_folder+"/journal_with_entry", "md", "", "")
	assert.Empty(t, ctx.Files())

	ctx = NewContext(test_folder+"/journal_with_various_entries", "md", "", "")
	assert.Equal(t, []string{test_folder + "/journal_with_various_entries/" + "1984-jan-25.md"}, ctx.Files())

	ctx = NewContext(test_folder+"/journal_zero", "", "", "")
	assert.Empty(t, ctx.Files())
}

func TestNewContext(t *testing.T) {
	ctx := NewContext("", "", "", "")
	homeDir, _ := user.Current()

	assert.Equal(t, homeDir.HomeDir+"/Journal/", ctx.Directory())
	assert.Equal(t, "txt", ctx.FileExtension())
	assert.Equal(t, "nano", ctx.Editor())
	assert.Equal(t, "", ctx.Template())
}

func TestNewContextWithEditor(t *testing.T) {
	assert.Equal(t, "vim", NewContext("", "", "vim", "").Editor())
}

func TestNewContextWithExtension(t *testing.T) {
	assert.Equal(t, "md", NewContext("", "md", "", "").FileExtension())
}

func TestNewContextWithDirectory(t *testing.T) {
	homeDir, _ := user.Current()
	ctx := NewContext("~/Journal/test", "", "", "")

	assert.Equal(t, homeDir.HomeDir+"/Journal/test/", ctx.Directory())
}

func TestNewContextWithTemplate(t *testing.T) {
	homeDir, _ := user.Current()
	ctx := NewContext("", "", "", "~/Journal/template.md")

	assert.Equal(t, homeDir.HomeDir+"/Journal/template.md", ctx.Template())
}
