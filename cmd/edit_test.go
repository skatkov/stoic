package stoic

import (
	"testing"
	"time"

	s "github.com/skatkov/stoic/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewEditCommand(t *testing.T) {
	ctx := s.NewContext("", "", "", "")
	command := NewEditCommand(ctx, "yesterday")

	assert.Equal(t, time.Now().AddDate(0, 0, -1), command.Date())
}
