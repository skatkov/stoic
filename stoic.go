package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

const DEFAULT_EDITOR = "nano"
const DEFAULT_DIRECTORY = "~/Journal/"
const DEFAULT_EXTENSION = "txt"

func main() {
	dir := directory()

	err := createDirectoryIfMissing(dir)
	if err != nil {
		return
	}

	openInEditor(dir + timeToFilename(time.Now()) + "." + DEFAULT_EXTENSION)
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

func timeToFilename(timestamp time.Time) string {
	year, month, day := timestamp.Date()

	return fmt.Sprintf("%d-%s-%02d", year, strings.ToLower(month.String()[:3]), day)
}

func openInEditor(filename string) error {
	var editor = os.Getenv("EDITOR")
	if editor == "" {
		editor = DEFAULT_EDITOR
	}
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func directory() string {
	directory := os.Getenv("STOIC_DIR")

	if directory == "" {
		directory, _ = homedir.Expand(DEFAULT_DIRECTORY)
	} else {
		directory, _ = homedir.Expand(directory)
	}

	return directory + "/"
}
