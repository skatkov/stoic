package stoic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testNewEntry(t *testing.T) {
	tm, _ := time.Parse("2006-Jan-02", "2020-Jan-01")
	e := NewEntry(tm, "~/Journal/")

	assert.Equal(t, "2020-Jan-01.txt", e.Filename())
	assert.Equal(t, "/", e.Filepath())
}
