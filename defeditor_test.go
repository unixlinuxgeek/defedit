package defeditor

import (
	"testing"
)

func TestOpen(t *testing.T) {
	s, err := Open()
	if err != nil {
		t.Errorf("open_editor: %s\n", err)
	}
	t.Logf("%s\n", s)
}
