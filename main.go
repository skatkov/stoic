package main

import (
	"flag"
	"os"
	"time"

	cmd "github.com/skatkov/stoic/cmd"
	stoic "github.com/skatkov/stoic/internal"
)

var BinaryVersion string   // Set via build flag
var BinaryBuildHash string // Set via build flag

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

	if *aboutFlag {
		cmd.NewAboutCommand(BinaryVersion, BinaryBuildHash).Run()
	} else if *listFlag {
		cmd.NewListCommand(ctx).Run()
	} else if *editFlag != "" {
		cmd.NewEditCommand(ctx, *editFlag).Run()
	} else if *quoteFlag {
		cmd.NewQuoteCommand().Run()
	} else {
		_ = ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
	}
}
