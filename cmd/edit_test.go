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
	y, m, d := time.Now().AddDate(0, 0, -1).Date()

	assert.Equal(t, time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location()), command.Date())
}
