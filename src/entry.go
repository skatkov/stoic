package stoic

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FILE_TEMPLATE = "2006-Jan-02"
const DEFAULT_EXTENSION = "txt"

type Entry interface {
	Filename() string
	Filepath() string
}

func NewEntry(time time.Time, dir string) Entry {
	return &entry{
		directory:  dir,
		created_at: time,
	}
}

type entry struct {
	directory  string
	created_at time.Time
}

func (e *entry) Filename() string {
	return strings.ToLower(fmt.Sprintf("%s.%s", time.Now().Format(FILE_TEMPLATE), fileExtension()))
}

func (e *entry) Filepath() string {
	return e.directory + e.Filename()
}

func fileExtension() string {
	extension := os.Getenv("STOIC_EXT")

	if extension == "" {
		extension = DEFAULT_EXTENSION
	}

	return extension
}
