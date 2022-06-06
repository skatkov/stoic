package stoic

import (
	"fmt"
	"strings"
	"time"
)

const FILE_TEMPLATE = "2006-jan-02"

type Entry interface {
	Filepath() string
}
type entry struct {
	filepath string
}

func NewEntry(ctx Context, time time.Time) Entry {
	return &entry{
		filepath: ctx.Directory() + strings.ToLower(fmt.Sprintf("%s.%s", time.Format(FILE_TEMPLATE), ctx.FileExtension())),
	}
}

func (e *entry) Filepath() string { return e.filepath }
