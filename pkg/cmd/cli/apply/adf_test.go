package apply

import (
	"testing"
)

func TestNewCmdApplyAzureRm(t *testing.T) {
	cmd := NewCmdApplyADF()

	if cmd.Use != "adf" {
		t.Errorf("Use is not correct")
	}
}
