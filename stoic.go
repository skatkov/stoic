package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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
	flag.Parse()

	if *aboutFlag {
		about_message := fmt.Sprintf("Version: %s \n", BinaryVersion)
		about_message += fmt.Sprintf("Build Hash: %s \n", BinaryBuildHash)

		fmt.Println(about_message)
		return //We're done here
	}

	_ = ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
}
