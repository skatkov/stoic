package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	stoic "github.com/skatkov/stoic/src"
)

var BinaryVersion string   // Set via build flag
var BinaryBuildHash string // Set via build flag

const DEFAULT_EDITOR = "nano"
const DEFAULT_DIRECTORY = "~/Journal/"

func main() {
	if about := about(); about != "" {
		fmt.Println(about)
		return
	}

	dir := directory()
	err := createDirectoryIfMissing(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	entry := stoic.NewEntry(time.Now(), dir)
	template, _ := homedir.Expand(os.Getenv("STOIC_TEMPLATE"))

	if template != "" && !fileExists(entry.Filepath()) {
		createFileFromTemplate(entry.Filepath(), template)
	}

	openInEditor(entry.Filepath())
}

func about() string {
	about := flag.Bool("about", false, "display about info")
	flag.Parse()

	if *about {
		return fmt.Sprintf("Version: %s", BinaryVersion) + "\n" + fmt.Sprintf("Build Hash: %s", BinaryBuildHash)
	}

	return ""
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n"), nil
}

func createFileFromTemplate(filename string, template_path string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	template, _ := readFile(template_path)

	_, err = file.WriteString(template)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
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
