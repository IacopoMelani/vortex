package cmd

//DeployCmd - Defines command to deploy current host as Vortex node
type DeployCmd struct {
	StandardCmd
}

// NewDeployCmd - Return a new instance of DeployCmd
func NewDeployCmd() *DeployCmd {
	return &DeployCmd{
		StandardCmd{
			Name:        CommandDeployNode,
			Description: "Deploy current host as node of Vortex network",
			Usage:       "vortex deploy",
			Flags:       []Flag{},
		},
	}
}

// CommandExec - Execs the command
func (j DeployCmd) CommandExec() error {

	println("deploy")

	return nil
}
