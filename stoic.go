package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
		p := tea.NewProgram(entryList{
			choices: []Entry{NewEntry(ctx, time.Now())},
		})
		if err := p.Start(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		return // We're done here.
	}

	if *about {
		fmt.Println(fmt.Sprintf("Version: %s", BinaryVersion) + "\n" + fmt.Sprintf("Build Hash: %s", BinaryBuildHash))
		return // We're done here.
	}

	_ = ctx.OpenInEditor(NewEntry(ctx, time.Now()))
}

type entryList struct {
	choices []Entry
}

func (m entryList) Init() tea.Cmd {
	return nil
}

func (m entryList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Was ia a keypress?
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m entryList) View() string {
	s := "Listing entries...\n\n"

	s += "\nPreq q to quit. \n"
	return s
}
