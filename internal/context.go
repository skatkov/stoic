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
	EnsureEntry(entry Entry) error
	Append(entry Entry, text string) error
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
	if err := ctx.EnsureEntry(entry); err != nil {
		return err
	}

	cmd := exec.Command(ctx.editor, entry.Filepath())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// EnsureEntry makes sure the journal directory exists and the entry file
// is created from template when configured.
func (ctx *context) EnsureEntry(entry Entry) error {
	if err := createDirectoryIfMissing(ctx.directory); err != nil {
		fmt.Println(err)
		return err
	}

	if ctx.Template() != "" && !fileExists(entry.Filepath()) {
		if err := createFileFromTemplate(entry.Filepath(), ctx.Template()); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if !fileExists(entry.Filepath()) {
		file, err := os.Create(entry.Filepath())
		if err != nil {
			fmt.Println(err)
			return err
		}
		_ = file.Close()
	}

	return nil
}

// Append adds text to the entry, ensuring a trailing newline separation.
func (ctx *context) Append(entry Entry, text string) error {
	if err := ctx.EnsureEntry(entry); err != nil {
		return err
	}

	f, err := os.OpenFile(entry.Filepath(), os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	separator, err := separatorForAppend(entry.Filepath())
	if err != nil {
		return err
	}

	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}

	_, err = f.WriteString(separator + text)
	return err
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

// separatorForAppend returns the separator to place before appended text.
// - Empty file: no separator.
// - Ends without newline: add "\n\n" (line break + one blank line).
// - Ends with one newline: add "\n" to create a blank line.
// - Ends with two or more newlines: add nothing.
func separatorForAppend(filepath string) (string, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return "", err
	}
	if info.Size() == 0 {
		return "", nil
	}

	trailing, err := trailingNewlines(filepath, info.Size())
	if err != nil {
		return "", err
	}

	switch {
	case trailing >= 2:
		return "", nil
	case trailing == 1:
		return "\n", nil
	default:
		return "\n\n", nil
	}
}

// trailingNewlines counts trailing '\n' characters up to the last two bytes.
func trailingNewlines(filepath string, size int64) (int, error) {
	readLen := int64(2)
	if size < readLen {
		readLen = size
	}

	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, readLen)
	_, err = file.ReadAt(buf, size-readLen)
	if err != nil {
		return 0, err
	}

	count := 0
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == '\n' {
			count++
		} else {
			break
		}
	}

	return count, nil
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
