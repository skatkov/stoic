package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const DEFAULT_EDITOR = "nano"

var extension = "txt"
var directory = "~/MEGAsync/journal/"

func main() {
	openInEditor(directory + timeToFilename(time.Now()) + "." + extension)
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
