package main

import (
	"flag"
	"fmt"
	"os"
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
	default:
		err := ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
		if err != nil {
			fmt.Println(err)
		}
	}
}
