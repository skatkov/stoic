package stoic

import (
	"fmt"
	"os"
	"time"

	stoic "github.com/skatkov/stoic/internal"
)

type EditCommand interface {
	Run()
}

type editCommand struct {
	ctx      stoic.Context
	filepath string
}

func NewEditCommand(ctx stoic.Context, filepath string) EditCommand {
	return &editCommand{
		ctx:      ctx,
		filepath: filepath,
	}
}

func (e *editCommand) Run() {
	entry := stoic.NewEntry(e.ctx, time.Now())
	err := e.ctx.OpenInEditor(entry)

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
