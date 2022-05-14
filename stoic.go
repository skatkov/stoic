package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

//https://github.com/jotaen/klog/blob/7b5503c7e39a94f6bdf5213c43e41e8ec5422064/src/app/context.go#L352

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

func tryCommands(commands []string, additionalArgs string) bool {
	for _, command := range commands {
		args := strings.Split(command, " ")
		args = append(args, additionalArgs)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err == nil {
			return true
		}
	}

	return false
}
