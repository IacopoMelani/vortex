package cmd

import (
	"fmt"
	"os"
)

// MARK: Command & Flag interface

// Command - Defines a generic interface for a command
type Command interface {
	// CommandExec - Execs the command
	CommandExec()
	// GetCommandDescription - Returns a description of command
	GetCommandDescription() string
	// GetCommandName - Returns the command name
	GetCommandName() string
	// GetCommandFlags - Returns a slice of string
	GetCommandFlags() []Flag
	// GetCommandUsage - Returns command usages
	GetCommandUsage() string
}

// Flag - Defines a generic interface for flags command
type Flag interface {
	// GetFlagName - Returns the flag's name
	GetFlagName() string
	// GetFlagDescription - Returns the flag description
	GetFlagDescription() string
	// GetFlagShortVersion - Returns the short version of the flag, empty flag if not defined
	GetFlagShortVersion() string
	// GetFlagVerboseVersion - Returns the verbose version of the flag, empty flag if not defined
	GetFlagVerboseVersion() string
	// GetFlagUsage - Returns flag usages
	GetFlagUsage() string
}

var availableCommands []Command

func init() {

	availableCommands = make([]Command, 0)

	availableCommands = []Command{
		*NewJoinTokenCmd(),
	}
}

// Parse - Parse the args in the command line
func Parse() error {

	var selectedCommand Command = nil

	for _, arg := range os.Args {

		if selectedCommand != nil {
			break
		}

		for _, command := range availableCommands {

			if command.GetCommandName() == arg {
				selectedCommand = command
				break
			}
		}
	}

	if selectedCommand == nil {
		fmt.Println("No commands specified")
		os.Exit(1)
	}

	selectedCommand.CommandExec()

	return nil
}

// MARK: StandardCmd & Command implementation

// StandardCmd - Defines the generic struct for Command implementation
type StandardCmd struct {
	Name        string
	Description string
	Flags       []Flag
	Usage       string
}

// GetCommandDescription - Returns a description of command
func (s StandardCmd) GetCommandDescription() string {
	return s.Description
}

// GetCommandName - Returns the command name
func (s StandardCmd) GetCommandName() string {
	return s.Name
}

// GetCommandFlags - Returns a slice of string
func (s StandardCmd) GetCommandFlags() []Flag {
	return s.Flags
}

// GetCommandUsage - Returns command usages
func (s StandardCmd) GetCommandUsage() string {
	return s.Usage
}

// MARK: StandardCmdFlag & Flag implementation

// StandardCmdFlag - Defines the generic struct for Flag implementation
type StandardCmdFlag struct {
	Name           string
	Description    string
	ShortVersion   string
	VerboseVersion string
	Usage          string
}

// GetFlagDescription - Returns the flag description
func (s StandardCmdFlag) GetFlagDescription() string {
	return s.Description
}

// GetFlagName - Returns the flag's name
func (s StandardCmdFlag) GetFlagName() string {
	return s.Name
}

// GetFlagShortVersion - Returns the short version of the flag, empty flag if not defined
func (s StandardCmdFlag) GetFlagShortVersion() string {
	return s.ShortVersion
}

// GetFlagVerboseVersion - Returns the verbose version of the flag, empty flag if not defined
func (s StandardCmdFlag) GetFlagVerboseVersion() string {
	return s.VerboseVersion
}

// GetFlagUsage - Returns flag usages
func (s StandardCmdFlag) GetFlagUsage() string {
	return s.Usage
}

// MARK: Info commands consts

const (
	CommandBase = "vortex"

	CommandDeployNode       = "deploy"
	CommndGenerateJoinToken = "join-token"
	CommandJoinToNode       = "join"
)

// MARK: Info commands Exported

// GetCommandDeployNode - Returns the command to deploy node on the network
func GetCommandDeployNode() string {
	return fmt.Sprintf("%s %s", CommandBase, CommandDeployNode)
}

// GetCommandGenerateJoinToken - Returns the command to generate a join token
func GetCommandGenerateJoinToken() string {
	return fmt.Sprintf("%s %s", CommandBase, CommndGenerateJoinToken)
}

// GetCommandJoinToNode - Returns the command to join to a node with join token
func GetCommandJoinToNode() string {
	return fmt.Sprintf("%s %s", CommandBase, CommandJoinToNode)
}

// MARK: JoinTokenCmd, JoinTokenCmdFlag & constructors

// JoinTokenCmd - Defines the command for generating join token
type JoinTokenCmd struct {
	StandardCmd
}

// NewJoinTokenCmd - Returns a new instance of JoinTokenCmd
func NewJoinTokenCmd() *JoinTokenCmd {
	return &JoinTokenCmd{
		StandardCmd: StandardCmd{
			Name:        "join-token",
			Description: "Generate a single-use join token to the vortex network",
			Usage:       "Go and use",
			Flags:       make([]Flag, 0),
		},
	}
}

func (j JoinTokenCmd) CommandExec() {
	println("Good")
}
