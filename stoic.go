package main

import (
	"flag"
	"fmt"
	"log"
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
	editFlag := flag.String("edit", "", "edit journal by filename")
	flag.Parse()

	if *aboutFlag {
		about_message := fmt.Sprintf("Version: %s \n", BinaryVersion)
		about_message += fmt.Sprintf("Build Hash: %s \n", BinaryBuildHash)

		fmt.Println(about_message)
		return //We're done here
	}

	if *editFlag != "" {
		entry, err := stoic.NewEntryFromString(ctx, *editFlag)
		if err != nil {
			log.Fatalf("invalid journal filename: %s", err)
		}
		_ = ctx.OpenInEditor(entry)

		return
	}

	_ = ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
}
