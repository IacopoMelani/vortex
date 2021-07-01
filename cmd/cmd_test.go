package cmd

import (
	"os"
	"testing"
)

func TestCmd(t *testing.T) {

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{CommandBase}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}
}
