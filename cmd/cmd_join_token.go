package cmd

import (
	"fmt"
	"os"

	"github.com/IacopoMelani/vortex/network"
)

const (
	JoinTokenCmdFlagHelp = "Help"
	JoinTokenCmdFlagHost = "Host"
)

// JoinTokenCmd - Defines the command for generating join token
type JoinTokenCmd struct {
	StandardCmd
}

// NewJoinTokenCmd - Returns a new instance of JoinTokenCmd
func NewJoinTokenCmd() *JoinTokenCmd {
	return &JoinTokenCmd{
		StandardCmd: StandardCmd{
			Name:        "join-token",
			Description: "Generates a single-use join token to the vortex network",
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
					Usage:          "join-token -H <host> | join-token --host=<host>",
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

	_, okHelp := j.IsCommandFlagUsed(JoinTokenCmdFlagHelp)

	if okHelp {

		ShowCommandHelp(j, true)
		return
	}

	jtConfig := network.JoinTokenConfig{}

	hostFlag, ok := j.IsCommandFlagUsed(JoinTokenCmdFlagHost)

	if ok {

		host := hostFlag.GetFlagValue()
		if host == "" {
			ShowError("No host provided!")

			hostFlag, ok := j.GetCommandFlagByName(JoinTokenCmdFlagHost)
			if ok {
				ShowFlagHelp(hostFlag, true)
			}
			os.Exit(1)
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
