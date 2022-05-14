package main

import (
	"fmt"
	"os"
	"os/exec"
)

var editor = "/usr/bin/nano"
var filename = "11-may-2022"
var extension = "txt"
var directory = "~/MEGAsync/journal/" + filename + "." + extension

// https://golangdocs.com/system-programming-in-go-3

func main() {
	fmt.Println(
		"Directory: " + directory,
	)
	openInEditor(directory)
}

func openInEditor(filename string) error {
	var editor = os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
