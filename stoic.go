package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	ctx := stoic.NewContext(
		os.Getenv("STOIC_DIR"),
		os.Getenv("STOIC_EXT"),
		os.Getenv("EDITOR"),
	)

	if about := about(); about != "" {
		fmt.Println(about)
		return
	}

	entry := stoic.NewEntry(ctx, time.Now())
	template, _ := homedir.Expand(os.Getenv("STOIC_TEMPLATE"))

	if template != "" && !fileExists(entry.Filepath()) {
		createFileFromTemplate(entry.Filepath(), template)
	}

	ctx.OpenInEditor(entry.Filepath())
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
