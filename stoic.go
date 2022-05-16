package main

import (
	"bufio"
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
		fmt.Println(err)
		return
	}

	openInEditor(dir + timeToFilename(time.Now()))
}

func fileExtension() string {
	extension := os.Getenv("STOIC_EXT")

	if extension == "" {
		extension = DEFAULT_EXTENSION
	}

	return extension
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

func timeToFilename(timestamp time.Time) string {
	year, month, day := timestamp.Date()
	//FIXME: This should be possible to format correctly, without modification on month variable
	filename := fmt.Sprintf("%d-%s-%02d", year, strings.ToLower(month.String()[:3]), day)

	return filename + "." + fileExtension()
}

func openInEditor(filename string) error {
	template, _ := homedir.Expand(os.Getenv("STOIC_TEMPLATE"))

	if template != "" && !fileExists(filename) {
		createFileFromTemplate(filename, template)
	}

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
