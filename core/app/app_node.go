package app

import (
	"sync"

	"github.com/IacopoMelani/vortex/core/network"
)

// MARK: AppNode, consts & constructors

const (
	// current Vortex node version
	VortexNodeVersion = "0.0.0"
)

// AppNode - Defines the Application for Vortex Network
type AppNode struct {
	AppStandard
	sync.RWMutex
	node *network.Node
}

// NewAppNode - Returns an instance of Application for Vortex Network
func NewAppNode(name string) *AppNode {

	app := NewApp(name, VortexNodeVersion, VortexModeNode)

	return &AppNode{
		AppStandard: *app,
		node:        nil,
	}
}

// MARK: AppNode Application implementation

// ID - Returns the Application ID
func (an *AppNode) ID() string {
	an.RLock()
	defer an.RUnlock()
	return an.id
}

// Mode - Returns the Application Mode
func (an *AppNode) Mode() string {
	an.RLock()
	defer an.RUnlock()
	return an.mode
}

// Name - Returns the Application name
func (an *AppNode) Name() string {
	an.RLock()
	defer an.RUnlock()
	return an.name
}

// Version - Returns the Application version
func (an *AppNode) Version() string {
	an.RLock()
	defer an.RUnlock()
	return an.version
}
