package cmd_test

import (
	"bytes"
	"fly-pokeapi/src/cmd"
	"testing"
)

func TestRootCmd_Execute(t *testing.T) {
	// Create the root command
	cmd := cmd.RootCommand()

	// Redirect stdout to a buffer to capture the output
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SetErr(out)

	// Execute the root command
	err := cmd.Execute()

	// Check for any execution error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the output
	expectedOutput := "Welcome to FlyPoke!!\n"
	if out.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, out.String())
	}
}
