package pony

import (
	"testing"
)

func TestNewCmdRoot(t *testing.T) {
	cmd := NewRootCmd("0.0.0", "commitid", "date")

	if cmd.Use != "pony" {
		t.Errorf("Use is not correct")
	}
}
