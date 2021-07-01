package cmd

import (
	"fmt"
	"os"
	"strings"
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
	// FlagIsPresent - Returns if the flag is passed
	FlagIsPresent() bool
	// FlagNeedValue - Returns true if the flag need a value in next os.Args
	FlagNeedValue() bool
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
	// GetFlagValue - Return the flag values passed
	GetFlagValue() string
	// SetFlagPresent - Set the value of Present for the flag
	SetFlagPresent(value bool)
	// SetFlagValue - Set the value passsed for the flag
	SetFlagValue(value string)
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

	if len(os.Args) > 2 {

		for index, argFlag := range os.Args[2:] {

			for _, flag := range selectedCommand.GetCommandFlags() {

				if argFlag == flag.GetFlagShortVersion() {

					flag.SetFlagPresent(true)

					if flag.FlagNeedValue() && len(os.Args) > index+2 && flag.GetFlagValue() == "" {

						flag.SetFlagValue(os.Args[index+2+1])

					}

				} else if strings.Contains(argFlag, flag.GetFlagVerboseVersion()) {

					flag.SetFlagPresent(true)

					if flag.FlagNeedValue() && flag.GetFlagValue() == "" {

						splittedArgFlag := strings.Split(argFlag, "=")

						if len(splittedArgFlag) > 1 {
							flag.SetFlagValue(splittedArgFlag[1])
						}
					}
				}
			}
		}

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
	Present        bool
	NeedValue      bool
	Value          string
}

// FlagIsPresent - Returns if the flag is passed
func (s StandardCmdFlag) FlagIsPresent() bool {
	return s.Present
}

// FlagNeedValue - Returns true if the flag need a value in next os.Args
func (s StandardCmdFlag) FlagNeedValue() bool {
	return s.NeedValue
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

// GetFlagValue - Return the flag values passed
func (s StandardCmdFlag) GetFlagValue() string {
	return s.Value
}

// SetFlagPresent - Set the value of Present for the flag
func (s *StandardCmdFlag) SetFlagPresent(value bool) {
	s.Present = value
}

// SetFlagValue - Set the value passsed for the flag
func (s *StandardCmdFlag) SetFlagValue(value string) {
	s.Value = value
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
			Flags:       []Flag{&StandardCmdFlag{Name: "sASSO", ShortVersion: "-s", VerboseVersion: "--sasso", NeedValue: true}},
		},
	}
}

func (j JoinTokenCmd) CommandExec() {
	println("Join token")
}
