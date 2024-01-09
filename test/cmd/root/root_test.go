package root_test

import (
	"bytes"
	"testing"

	"github.com/JustinJohnsonK/macu-cli/cmd/root"
)

func TestRootCmd(t *testing.T) {
	cmd := root.RootCmd()

	var testOut bytes.Buffer
	cmd.SetOut(&testOut)

	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expectedOut := "My Macu is a good girl"
	if testOut.String() != expectedOut {
		t.Errorf("Expected Output: %q, Actual Output: %s", expectedOut, testOut.String())
	}
}
