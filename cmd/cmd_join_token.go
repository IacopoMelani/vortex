package cmd

import (
	"fmt"

	"github.com/IacopoMelani/vortex/core/network"
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
			Name:        CommndGenerateJoinToken,
			Description: "Generates a single-use join token to the vortex network",
			Usage:       "vortex join-token",
			Flags: []Flag{
				&StandardCmdFlag{
					Name:           JoinTokenCmdFlagHelp,
					Description:    "Show this message",
					Usage:          "join-token -h | join-token --help",
					ShortVersion:   "-h",
					VerboseVersion: "--help",
					Present:        false,
					NeedValue:      false,
				},
				&StandardCmdFlag{
					Name:           JoinTokenCmdFlagHost,
					Description:    "Used for specify the host for join token",
					Usage:          "join-token -H <host> | join-token --host=<host>",
					VerboseVersion: "--host",
					ShortVersion:   "-H",
					Present:        false,
					NeedValue:      true,
				},
			},
		},
	}
}

// CommandExec - Execs the command
func (j JoinTokenCmd) CommandExec() error {

	_, okHelp := j.IsCommandFlagUsed(JoinTokenCmdFlagHelp)

	if okHelp {

		ShowCommandHelp(j, true)
		return nil
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
			return nil
		}

		jtConfig.Host = host
	}

	joinToken, err := network.NewJoinTokenWithConfig(jtConfig)
	if err != nil {
		return err
	}

	fmt.Println(j.JoinCommandSample(joinToken))

	return nil
}

// JoinCommand - Returns the complete command to join a node
func (j JoinTokenCmd) JoinCommandSample(jt *network.JoinToken) string {

	return fmt.Sprintf("\n%s --host=%s --token=%s\n", GetCommandJoinToNode(), jt.Host(), jt.Value())
}
