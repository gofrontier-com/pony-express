package validate

import (
	"testing"
)

func TestNewCmdValidateAzureRm(t *testing.T) {
	cmd := NewCmdValidateADF()

	if cmd.Use != "adf" {
		t.Errorf("Use is not correct")
	}
}
