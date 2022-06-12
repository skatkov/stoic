package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		about_message := fmt.Sprintf("Version: %s \n", BinaryVersion)
		about_message += fmt.Sprintf("Build Hash: %s \n", BinaryBuildHash)

		fmt.Println(about_message)
		return //We're done here
	} else if *listFlag {
		var items []list.Item
		files := ctx.Files()

		for _, file := range files {
			items = append(items, item{
				title: file,
				desc:  "test",
			})
		}

		m := model{
			list:    list.New(items, list.NewDefaultDelegate(), 0, 0),
			context: ctx,
		}
		m.list.Title = "Journal Entries"

		p := tea.NewProgram(m, tea.WithAltScreen())

		if err := p.Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	} else {
		_ = ctx.OpenInEditor(stoic.NewEntry(ctx, time.Now()))
	}
}

// --- List feature

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    list.Model
	context stoic.Context
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else if msg.String() == "enter" {
			fmt.Println("enter was pressed")
		} else if msg.String() == " " {
			fmt.Println("space was pressed")
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
