package main

import (
	"flag"
	"fmt"
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

	defaultTime := boolFromEnv("STOIC_APPEND_TIME")
	defaultBullet := boolFromEnv("STOIC_APPEND_BULLET")
	timeFlag := flag.Bool("time", defaultTime, "prefix appended text with current time (HH:MM)")
	flag.BoolVar(timeFlag, "t", defaultTime, "prefix appended text with current time (HH:MM)")
	bulletFlag := flag.Bool("bullet", defaultBullet, "prefix appended text with '- '")
	flag.BoolVar(bulletFlag, "b", defaultBullet, "prefix appended text with '- '")
	flag.Parse()

	switch {
	case *aboutFlag:
		cmd.NewAboutCommand(version, commit, date).Run()
	case *listFlag:
		cmd.NewListCommand(ctx).Run()
	case *editFlag != "":
		cmd.NewEditCommand(ctx, *editFlag).Run()
	case *quoteFlag:
		cmd.NewQuoteCommand().Run()
	case len(flag.Args()) > 0:
		text := strings.Join(flag.Args(), " ")
		cmd.NewAppendCommand(ctx, time.Now(), text, *timeFlag, *bulletFlag).Run()
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
