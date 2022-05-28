package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	stoic "github.com/skatkov/stoic/src"
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

	if about := about(); about != "" {
		fmt.Println(about)
		return
	}

	entry := stoic.NewEntry(ctx, time.Now())

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
