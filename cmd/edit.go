package stoic

import (
	"fmt"
	"os"
	"time"

	stoic "github.com/skatkov/stoic/internal"
	naturaldate "github.com/tj/go-naturaldate"
)

type EditCommand interface {
	Run()
	Date() time.Time
}

type editCommand struct {
	ctx   stoic.Context
	value string
	date  time.Time
}

func NewEditCommand(ctx stoic.Context, value string) EditCommand {
	date, err := naturaldate.Parse(value, time.Now())
	if err != nil {
		fmt.Println("Error parsing date:", err)
		os.Exit(1)
	}

	return &editCommand{
		ctx:   ctx,
		value: value,
		date:  date,
	}
}

func (e *editCommand) Date() time.Time {
	return e.date
}

func (e *editCommand) Run() {
	entry := stoic.NewEntry(e.ctx, e.date)
	err := e.ctx.OpenInEditor(entry)
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
