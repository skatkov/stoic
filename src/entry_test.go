package stoic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	tm, _ := time.Parse("2006-Jan-02", "2020-Jan-01")
	ctx := NewContext("", "", "")
	e := NewEntry(ctx, tm)

	assert.Equal(t, "~/Journal/2020-jan-01.txt", e.Filepath())
}
