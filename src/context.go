package stoic

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

const DEFAULT_EDITOR = "nano"
const DEFAULT_DIRECTORY = "~/Journal/"
const DEFAULT_EXTENSION = "txt"

type Context interface {
	Directory() string
	FileExtension() string
	Template() string
	OpenInEditor(filepath string) error
}

type context struct {
	directory     string
	fileExtension string
	editor        string
	template      string
}

func NewContext(homeDir string, fileExtension string, editor string, template string) Context {
	directory := directory(homeDir)
	err := createDirectoryIfMissing(directory)
	if err != nil {
		fmt.Println(err)
	}

	if fileExtension == "" {
		fileExtension = DEFAULT_EXTENSION
	}

	template, _ = homedir.Expand(template)

	if editor == "" {
		editor = DEFAULT_EDITOR
	}

	return &context{
		directory:     directory,
		fileExtension: fileExtension,
		editor:        editor,
		template:      template,
	}
}

func (ctx *context) Directory() string {
	return ctx.directory
}

func (ctx *context) FileExtension() string {
	return ctx.fileExtension
}

func (ctx *context) OpenInEditor(filepath string) error {
	if ctx.Template() != "" && !fileExists(filepath) {
		createFileFromTemplate(filepath, ctx.Template())
	}

	cmd := exec.Command(ctx.editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}

	return true
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

func (ctx *context) Template() string {
	return ctx.template
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
