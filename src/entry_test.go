package stoic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	tm, _ := time.Parse("2006-Jan-02", "2020-Jan-01")
	e := NewEntry(tm, "~/Journal/")

	assert.Equal(t, "2020-jan-01.md", e.Filename())
	assert.Equal(t, "~/Journal/2020-jan-01.md", e.Filepath())
}
