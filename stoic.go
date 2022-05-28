package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	. "github.com/skatkov/stoic/src"
)

var BinaryVersion string   // Set via build flag
var BinaryBuildHash string // Set via build flag

const defaultWidth = 20
const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

func main() {
	ctx := NewContext(
		os.Getenv("STOIC_DIR"),
		os.Getenv("STOIC_EXT"),
		os.Getenv("EDITOR"),
		os.Getenv("STOIC_TEMPLATE"),
	)

	listFlag := flag.Bool("list", false, "list all entries")
	aboutFlag := flag.Bool("about", false, "display about info")

	flag.Parse()

	if *listFlag {
		items := []list.Item{
			item("Ramen"),
			item("Tomato Soup"),
			item("Hamburgers"),
			item("Cheeseburgers"),
			item("Currywurst"),
			item("Okonomiyaki"),
			item("Pasta"),
			item("Fillet Mignon"),
			item("Caviar"),
			item("Just Wine"),
		}

		l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
		l.Title = "Which entry are you interested in?"
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.Styles.Title = titleStyle
		l.Styles.PaginationStyle = paginationStyle
		l.Styles.HelpStyle = helpStyle

		p := tea.NewProgram(entryList{
			list: l,
		})
		if err := p.Start(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		return // We're done here.
	}

	if *aboutFlag {
		fmt.Println(fmt.Sprintf("Version: %s", BinaryVersion) + "\n" + fmt.Sprintf("Build Hash: %s", BinaryBuildHash))
		return // We're done here.
	}

	_ = ctx.OpenInEditor(NewEntry(ctx, time.Now()))
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type entryList struct {
	list   list.Model
	items  []item
	choice string
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
	return "\n" + m.list.View()
}
