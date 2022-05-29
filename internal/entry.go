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
	directory  string
	created_at time.Time
	filename   string
}

func NewEntry(ctx Context, time time.Time) Entry {
	return &entry{
		filename:   strings.ToLower(fmt.Sprintf("%s.%s", time.Format(FILE_TEMPLATE), ctx.FileExtension())),
		directory:  ctx.Directory(),
		created_at: time,
	}
}

func (e *entry) Filepath() string {
	return e.directory + e.filename
}
