package stoic

import (
	"fmt"
	"os"
	"strings"
	"time"

	stoic "github.com/skatkov/stoic/internal"
)

type AppendCommand interface {
	Run()
}

type appendCommand struct {
	ctx       stoic.Context
	date      time.Time
	text      string
	useTime   bool
	useBullet bool
	entry     stoic.Entry
}

func NewAppendCommand(ctx stoic.Context, date time.Time, text string, useTime, useBullet bool) AppendCommand {
	entry := stoic.NewEntry(ctx, date)
	return &appendCommand{
		ctx:       ctx,
		date:      date,
		text:      strings.TrimSpace(text),
		useTime:   useTime,
		useBullet: useBullet,
		entry:     entry,
	}
}

func (a *appendCommand) Run() {
	if a.text == "" {
		fmt.Println("No text provided to append.")
		os.Exit(1)
	}

	formatted := formatAppendText(a.text, a.date, a.useTime, a.useBullet)

	if err := a.ctx.Append(a.entry, formatted); err != nil {
		fmt.Println("Error appending to entry:", err)
		os.Exit(1)
	}
}

func formatAppendText(text string, when time.Time, useTime, useBullet bool) string {
	prefixParts := []string{}

	if useBullet {
		prefixParts = append(prefixParts, "-")
	}

	if useTime {
		prefixParts = append(prefixParts, formatTimeHHMM(when))
	}

	if len(prefixParts) == 0 {
		return text
	}

	prefix := strings.Join(prefixParts, " ")
	if useTime {
		return fmt.Sprintf("%s — %s", prefix, text)
	}

	return fmt.Sprintf("%s %s", prefix, text)
}

func formatTimeHHMM(when time.Time) string {
	return fmt.Sprintf("%02d:%02d", when.Hour(), when.Minute())
}
