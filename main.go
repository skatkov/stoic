package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	cmd "github.com/skatkov/stoic/cmd"
	stoic "github.com/skatkov/stoic/internal"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	ctx := stoic.NewContext(
		os.Getenv("STOIC_DIR"),
		os.Getenv("STOIC_EXT"),
		os.Getenv("EDITOR"),
		os.Getenv("STOIC_TEMPLATE"),
	)

	aboutFlag := flag.Bool("about", false, "display about info")
	listFlag := flag.Bool("list", false, "list journal entries")
	quoteFlag := flag.Bool("quote", false, "random quote to inspire ongoing journaling habit")
	editFlag := flag.String("edit", "", "edit a journal entry")

	defaultBullet := boolFromEnv("STOIC_APPEND_BULLET")
	bulletFlag := flag.Bool("bullet", defaultBullet, "prefix appended text with '- '")
	flag.BoolVar(bulletFlag, "b", defaultBullet, "prefix appended text with '- '")
	flag.Parse()

	appendText, hasAppendText, err := resolveAppendInput(flag.Args(), os.Stdin)
	if err != nil {
		fmt.Println("Error reading piped input:", err)
		os.Exit(1)
	}

	switch {
	case *aboutFlag:
		cmd.NewAboutCommand(version, commit, date).Run()
	case *listFlag:
		cmd.NewListCommand(ctx).Run()
	case *editFlag != "":
		cmd.NewEditCommand(ctx, *editFlag).Run()
	case *quoteFlag:
		cmd.NewQuoteCommand().Run()
	case hasAppendText:
		cmd.NewAppendCommand(ctx, time.Now(), appendText, *bulletFlag).Run()
	default:
		err := ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func boolFromEnv(key string) bool {
	value := strings.ToLower(os.Getenv(key))
	switch value {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

// resolveAppendInput determines whether append text was provided either as
// positional arguments or via piped stdin. It returns the text, a boolean
// indicating presence, and any read error.
func resolveAppendInput(args []string, stdin *os.File) (string, bool, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), true, nil
	}

	if stdin == nil {
		return "", false, nil
	}

	info, err := stdin.Stat()
	if err != nil {
		return "", false, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return "", false, nil
	}

	data, err := io.ReadAll(stdin)
	if err != nil {
		return "", false, err
	}

	if len(data) == 0 {
		return "", false, nil
	}

	return string(data), true, nil
}
