package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	. "github.com/skatkov/stoic/src"
)

var BinaryVersion string   // Set via build flag
var BinaryBuildHash string // Set via build flag

func main() {
	ctx := NewContext(
		os.Getenv("STOIC_DIR"),
		os.Getenv("STOIC_EXT"),
		os.Getenv("EDITOR"),
		os.Getenv("STOIC_TEMPLATE"),
	)

	list := flag.Bool("list", false, "list all entries")
	about := flag.Bool("about", false, "display about info")

	flag.Parse()

	if *list {
		fmt.Println("Listing entries...")
		return
	}

	if *about {
		fmt.Println(fmt.Sprintf("Version: %s", BinaryVersion) + "\n" + fmt.Sprintf("Build Hash: %s", BinaryBuildHash))
		return
	}

	_ = ctx.OpenInEditor(NewEntry(ctx, time.Now()))
}
