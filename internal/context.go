package stoic

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	DEFAULT_EDITOR    = "nano"
	DEFAULT_DIRECTORY = "~/Journal/"
	DEFAULT_EXTENSION = "md"
)

type Context interface {
	Directory() string
	FileExtension() string
	Template() string
	Editor() string
	OpenInEditor(entry Entry) error
	Files() []string
}

type context struct {
	directory     string
	fileExtension string
	editor        string
	template      string
}

func NewContext(homeDir, fileExtension, editor, template string) Context {
	directory := expandDir(homeDir)

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

func (ctx *context) Directory() string     { return ctx.directory }
func (ctx *context) FileExtension() string { return ctx.fileExtension }
func (ctx *context) Editor() string        { return ctx.editor }
func (ctx *context) Template() string      { return ctx.template }

func (ctx *context) OpenInEditor(entry Entry) error {
	err := createDirectoryIfMissing(ctx.directory)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if ctx.Template() != "" && !fileExists(entry.Filepath()) {
		_ = createFileFromTemplate(entry.Filepath(), ctx.Template())
	}

	cmd := exec.Command(ctx.editor, entry.Filepath())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createFileFromTemplate(filename, template_path string) error {
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

func expandDir(directory string) string {
	if directory == "" {
		directory, _ = homedir.Expand(DEFAULT_DIRECTORY)
	} else {
		directory, _ = homedir.Expand(directory)
	}

	return directory + "/"
}

func (ctx context) Files() []string {
	files, _ := os.ReadDir(ctx.directory)

	var filenames []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ctx.fileExtension) {
			filename := strings.TrimSuffix(file.Name(), "."+ctx.fileExtension)

			_, err := time.Parse(FILE_TEMPLATE, filename)
			if err == nil {
				filenames = append(filenames, ctx.directory+file.Name())
			}
		}
	}

	return filenames
}

func createDirectoryIfMissing(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
	}

	return nil
}
