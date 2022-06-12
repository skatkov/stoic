package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	if *aboutFlag {
		about_message := fmt.Sprintf("Version: %s", BinaryVersion) + "\n"
		about_message += fmt.Sprintf("Build Hash: %s", BinaryBuildHash) + "\n"

		fmt.Println(about_message)
	} else if *listFlag {
		cmd := cmd.NewListCommand(ctx)
		cmd.Run()
	} else {
		_ = ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
	}
}
