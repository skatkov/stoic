package stoic

import (
	"fmt"
	"os"
	"strings"
	"time"

	stoic "github.com/skatkov/stoic/internal"
)

// AppendCommand appends text to a journal entry.
type AppendCommand interface {
	Run()
}

type appendCommand struct {
	ctx       stoic.Context
	text      string
	useBullet bool
	entry     stoic.Entry
}

// NewAppendCommand creates a command that appends text to the specified entry date.
func NewAppendCommand(ctx stoic.Context, date time.Time, text string, useBullet bool) AppendCommand {
	entry := stoic.NewEntry(ctx, date)
	return &appendCommand{
		ctx:       ctx,
		text:      strings.TrimSpace(text),
		useBullet: useBullet,
		entry:     entry,
	}
}

func (a *appendCommand) Run() {
	if a.text == "" {
		fmt.Println("No text provided to append.")
		os.Exit(1)
	}

	formatted := formatAppendText(a.text, a.useBullet)

	if err := a.ctx.Append(a.entry, formatted); err != nil {
		fmt.Println("Error appending to entry:", err)
		os.Exit(1)
	}
}

func formatAppendText(text string, useBullet bool) string {
	prefixParts := []string{}

	if useBullet {
		prefixParts = append(prefixParts, "-")
	}

	if len(prefixParts) == 0 {
		return text
	}

	prefix := strings.Join(prefixParts, " ")
	return fmt.Sprintf("%s %s", prefix, text)
}
