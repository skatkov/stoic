package stoic

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
