package stoic

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	stoic "github.com/skatkov/stoic/internal"
)

type ListCommand interface {
	Run()
}

type listCommand struct {
	ctx stoic.Context
}

func NewListCommand(ctx stoic.Context) ListCommand {
	return &listCommand{
		ctx: ctx,
	}
}

func (lc listCommand) Run() {
	var items []list.Item
	files := lc.ctx.Files()

	for _, file := range files {
		fileInfo, _ := os.Lstat(file)
		status := fmt.Sprintf("%s %s %s",
			fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			fileInfo.Mode().String(),
			ConvertBytesToSizeString(fileInfo.Size()))

		items = append(items, item{
			title: file,
			desc:  status,
		})
	}

	m := model{
		list:    list.New(items, list.NewDefaultDelegate(), 0, 0),
		context: lc.ctx,
	}
	m.list.Title = "Journal Entries"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil { // Updated from p.Start() to p.Run()
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

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
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case msg.String() == "enter":
			selectedItem, _ := m.list.SelectedItem().(item)

			_ = OpenFileInEditor(selectedItem.title, m.context)
			os.Exit(0)
		case msg.String() == " ":
			selectedItem, _ := m.list.SelectedItem().(item)

			_ = OpenFileInEditor(selectedItem.title, m.context)
			os.Exit(0)
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

func OpenFileInEditor(filepath string, ctx stoic.Context) error {
	cmd := exec.Command(ctx.Editor(), filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

const (
	thousand    = 1000
	ten         = 10
	fivePercent = 0.0499
)

// ConvertBytesToSizeString converts a byte count to a human readable string.
func ConvertBytesToSizeString(size int64) string {
	if size < thousand {
		return fmt.Sprintf("%dB", size)
	}

	suffix := []string{
		"K", // kilo
		"M", // mega
		"G", // giga
		"T", // tera
		"P", // peta
		"E", // exa
		"Z", // zeta
		"Y", // yotta
	}

	curr := float64(size) / thousand
	for _, s := range suffix {
		if curr < ten {
			return fmt.Sprintf("%.1f%s", curr-fivePercent, s)
		} else if curr < thousand {
			return fmt.Sprintf("%d%s", int(curr), s)
		}
		curr /= thousand
	}

	return ""
}
