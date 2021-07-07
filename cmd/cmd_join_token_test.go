package cmd

import (
	"os"
	"testing"
)

func TestCmdJoinToken(t *testing.T) {

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		appCLI.resetCommands()
	}()

	// vortex join-token

	os.Args = []string{CommandBase, CommndGenerateJoinToken}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}
}

func TestCmdJoinTokenHelp(t *testing.T) {

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		appCLI.resetCommands()
	}()

	// vortex join-token -h

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "-h"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	appCLI.resetCommands()

	// vortex join-token -h

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "--help"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}
}

func TestCmdJoinTokenHost(t *testing.T) {

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		appCLI.resetCommands()
	}()

	// vortex join-token -H

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "-H"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	appCLI.resetCommands()

	// vortex join-token -H <host>

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "-H", "127.0.0.1"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	appCLI.resetCommands()

	// vortex join-token --host

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "--host"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}

	appCLI.resetCommands()

	// vortex join-token ---host=<host>

	os.Args = []string{CommandBase, CommndGenerateJoinToken, "--host=127.0.0.1"}

	if err := Parse(); err != nil {
		t.Fatal(err)
	}
}
