package plan

import (
	"testing"
)

func TestNewCmdValidateAzureRm(t *testing.T) {
	cmd := NewCmdPlanADF()

	if cmd.Use != "adf" {
		t.Errorf("Use is not correct")
	}
}
