package main

import (
	"log"
	"os/exec"
)

var editor = "/usr/bin/nano"
var directory = "~/MEGAsync/journal/11-may-2022.txt"
var filename = ""

// https://golangdocs.com/system-programming-in-go-3

func main() {
	cmd := exec.Command(editor, directory)
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}