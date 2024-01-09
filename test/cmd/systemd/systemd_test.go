package systemd_test

import (
	"bytes"
	"testing"

	"github.com/JustinJohnsonK/macu-cli/cmd/systemd"
)

func TestSystemd(t *testing.T) {
	cmd := systemd.NewFile()

	// OutPut stream
	var testOut bytes.Buffer
	cmd.SetOut(&testOut)

	cmd.SetArgs([]string{"new-file"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expectedOut := "Created new-file and moved inside"
	if testOut.String() != expectedOut {
		t.Errorf("Expected Output: %q, Actual Output: %s", expectedOut, testOut.String())
	}

}
