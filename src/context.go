package stoic

import (
	"os"
	"os/exec"
)

const DEFAULT_EDITOR = "nano"
const DEFAULT_DIRECTORY = "~/Journal/"
const DEFAULT_EXTENSION = "txt"

type Context interface {
	OpenInEditor(filepath string) error
}

type context struct {
	directory     string
	fileExtension string
	editor        string
}

func NewContext(homeDir string, fileExtension string, editor string) Context {
	if homeDir == "" {
		homeDir = DEFAULT_DIRECTORY
	}

	if fileExtension == "" {
		fileExtension = DEFAULT_EXTENSION
	}

	if editor == "" {
		editor = DEFAULT_EDITOR
	}

	return &context{
		directory:     homeDir,
		fileExtension: fileExtension,
		editor:        editor,
	}
}

func (ctx *context) OpenInEditor(filepath string) error {
	cmd := exec.Command(ctx.editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
