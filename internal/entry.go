package stoic

import (
	"fmt"
	"strings"
	"time"
)

const FILE_TEMPLATE = "2006-Jan-02"

type Entry interface {
	Filepath() string
}
type entry struct {
	filepath string
}

func FindEntriesInFolder(ctx Context) []Entry {
	entries := make([]Entry, 0)
	for _, filename := range ctx.ListFiles() {
		if strings.HasSuffix(filename, ctx.FileExtension()) {
			entries = append(entries, NewEntry(ctx, time.Now()))
		}
	}
	return entries
}

func NewEntry(ctx Context, time time.Time) Entry {
	return &entry{
		filepath: ctx.Directory() + strings.ToLower(fmt.Sprintf("%s.%s", time.Format(FILE_TEMPLATE), ctx.FileExtension())),
	}
}

func (e *entry) Filepath() string { return e.filepath }
