package stoic

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mitchellh/go-homedir"
)

const DEFAULT_EDITOR = "nano"
const DEFAULT_DIRECTORY = "~/Journal/"
const DEFAULT_EXTENSION = "txt"

type Context interface {
	Directory() string
	FileExtension() string
	OpenInEditor(filepath string) error
}

type context struct {
	directory     string
	fileExtension string
	editor        string
}

func NewContext(homeDir string, fileExtension string, editor string) Context {
	directory := directory(homeDir)
	err := createDirectoryIfMissing(directory)
	if err != nil {
		fmt.Println(err)
	}

	if fileExtension == "" {
		fileExtension = DEFAULT_EXTENSION
	}

	if editor == "" {
		editor = DEFAULT_EDITOR
	}

	return &context{
		directory:     directory,
		fileExtension: fileExtension,
		editor:        editor,
	}
}

func (ctx *context) Directory() string {
	return ctx.directory
}

func (ctx *context) FileExtension() string {
	return ctx.fileExtension
}

func (ctx *context) OpenInEditor(filepath string) error {
	cmd := exec.Command(ctx.editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func directory(dir string) string {
	directory := dir

	if directory == "" {
		directory, _ = homedir.Expand(DEFAULT_DIRECTORY)
	} else {
		directory, _ = homedir.Expand(directory)
	}

	return directory + "/"
}

func createDirectoryIfMissing(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)

		if err != nil {
			return err
		}
	}

	return nil
}
