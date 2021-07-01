package app

import (
	"github.com/google/uuid"
)

// MARK: App Standard

// defines available Application Mode
const (
	VortexModeCLI      = "CLI"
	VortexModeNode     = "Node"
	VortexModeConsumer = "Consumer"
)

// Application - Defines a interface for an Application in Vortex system based on "mode", see const VortexMode*
type Application interface {
	// ID - Returns the Application ID
	ID() string
	// Mode - Returns the Application Mode
	Mode() string
	// Name - Returns the Application name
	Name() string
	// Version - Returns the Application version
	Version() string
}

// AppStandard - Defines standard field for an Application
type AppStandard struct {
	id      string
	name    string
	version string
	mode    string
}

// NewApp - Returns new AppStandard instance
func NewApp(name, version, mode string) *AppStandard {
	return &AppStandard{
		id:      uuid.NewString(),
		name:    name,
		version: version,
		mode:    mode,
	}
}
