package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/IacopoMelani/vortex/network"
)

// MARK: consts

const (
	VortexVersion = "0.0.1"
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
	// IsCommandFlagUsed - Returns The Flag interface if present
	IsCommandFlagUsed(name string) (Flag, bool)
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
		ShowHelp(true)
		os.Exit(1)
	}

	if len(os.Args) > 2 {

		for index, argFlag := range os.Args[2:] {

			for _, flag := range selectedCommand.GetCommandFlags() {

				if flag.GetFlagShortVersion() != "" && argFlag == flag.GetFlagShortVersion() {

					flag.SetFlagPresent(true)

					if flag.FlagNeedValue() && len(os.Args) > index+2 && flag.GetFlagValue() == "" {

						flag.SetFlagValue(os.Args[index+2+1])

					}

				} else if flag.GetFlagVerboseVersion() != "" && strings.Contains(argFlag, flag.GetFlagVerboseVersion()) {

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

// ShowBanner - Shows the banner
func ShowBanner() {
	b, err := ioutil.ReadFile("cmd/banner.txt")
	if err != nil {
		panic(err)
	}
	str := strings.Replace(string(b), "<VERSION>", VortexVersion, -1)
	fmt.Println(str)
}

// ShowCommandHelp -  Shows the help for current command
func ShowCommandHelp(command Command, withUsage bool) {

	fmt.Printf("%s help command list\n\n", command.GetCommandName())

	for _, flag := range command.GetCommandFlags() {
		ShowFlagHelp(flag, withUsage)
	}

	fmt.Printf("\n")
}

// ShowFlagHelp - Shows the help for current flag
func ShowFlagHelp(flag Flag, withUsage bool) {
	if withUsage {
		fmt.Printf("%s\t%s\n\t%s\n\n", flag.GetFlagName(), flag.GetFlagDescription(), flag.GetFlagUsage())
	} else {
		fmt.Printf("%s\t%s\n", flag.GetFlagName(), flag.GetFlagDescription())
	}
}

// ShowHelp - Shows the full help commands
func ShowHelp(withUsage bool) {
	ShowBanner()
	for _, command := range availableCommands {
		ShowCommandHelp(command, withUsage)
	}
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

// IsCommandFlagUsed - Returns The Flag interface if present
func (s StandardCmd) IsCommandFlagUsed(name string) (Flag, bool) {

	for _, flag := range s.Flags {

		if flag.GetFlagName() == name && flag.FlagIsPresent() {
			return flag, true
		}
	}

	return nil, false
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
			Flags: []Flag{
				&StandardCmdFlag{
					Name:           "Help",
					Description:    "Show this message",
					Usage:          "join-token -h | join-token --help",
					ShortVersion:   "-h",
					VerboseVersion: "--help",
				},
				&StandardCmdFlag{
					Name:           "Host",
					Description:    "Used for specify the host for join token",
					Usage:          "join-token -H <HOST> | join-token --host=<HOST>",
					VerboseVersion: "--host",
					ShortVersion:   "-H",
					NeedValue:      true,
				},
			},
		},
	}
}

// CommandExec - Execs the command
func (j JoinTokenCmd) CommandExec() {

	_, ok := j.IsCommandFlagUsed("Help")

	if ok {

		fmt.Printf("join-token help command list\n\n")

		for _, flag := range j.GetCommandFlags() {
			fmt.Printf("%s\t%s\n\t%s\n\n", flag.GetFlagName(), flag.GetFlagDescription(), flag.GetFlagUsage())
		}

		fmt.Printf("\n")
		return
	}

	jtConfig := network.JoinTokenConfig{}

	hostFlag, ok := j.IsCommandFlagUsed("Host")

	if ok {

		host := hostFlag.GetFlagValue()
		if host != "" {
			jtConfig.Host = host
		}
	}

	joinToken, err := network.NewJoinTokenWithConfig(jtConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(j.JoinCommandSample(joinToken))
}

// JoinCommand - Returns the complete command to join a node
func (j JoinTokenCmd) JoinCommandSample(jt *network.JoinToken) string {

	return fmt.Sprintf("\n%s --host=%s --token=%s\n", GetCommandJoinToNode(), jt.Host(), jt.Value())
}
