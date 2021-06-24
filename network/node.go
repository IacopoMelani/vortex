package network

import (
	"fmt"
	"os"
	"sync"

	"github.com/IacopoMelani/vortex-storage/utils"
	"github.com/google/uuid"
)

const (
	DefaultPort = ":6414"
)

// Node - Defines a node of vortex network
type Node struct {
	sync.RWMutex
	neighbors map[string]*Node
	host      string
	id        string
	name      string
	port      string
}

// NodeConfig - Defines a node config struct
type NodeConfig struct {
	Name string
	IP   string
	Port string
}

// NewNode - Returns a new instance of Node
func NewNode() (*Node, error) {

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	ip, err := utils.GetPrimaryIP()
	if err != nil {
		return nil, err
	}

	return &Node{
		id:        uuid.New().String(),
		name:      hostname,
		host:      ip.String(),
		port:      DefaultPort,
		neighbors: make(map[string]*Node),
	}, nil
}

// NewWithConfig - Return a new instance of Node with config passed
func NewWithConfig(config NodeConfig) (*Node, error) {

	node, err := NewNode()
	if err != nil {
		return nil, err
	}

	node.Lock()
	defer node.Unlock()

	if config.Name != "" {
		node.name = config.Name
	}
	if config.Port != "" {
		node.port = config.Port
	}
	if config.IP != "" {
		node.host = config.IP
	}

	return node, nil
}

// AddNeighbor - Add new node if the neighbors networks
func (n *Node) AddNeighbor(newNode *Node) error {

	n.Lock()
	defer n.Unlock()

	_, ok := n.neighbors[newNode.ID()]
	if ok {
		return NewNodeAlreadyNeighborError(newNode)
	}

	n.neighbors[newNode.ID()] = newNode

	return nil
}

// ID - Returns Node ID
func (n *Node) ID() string {
	n.RLock()
	defer n.RUnlock()
	return n.id
}

// Name - Returns node name
func (n *Node) Name() string {
	n.RLock()
	defer n.RUnlock()
	return n.name
}

// NodeAlreadyNeighborError - Defines error for
type NodeAlreadyNeighborError struct {
	nodeName string
}

// NewNodeAlreadyNeighborError - Returns a new instance of NodeAlreadyNeighborError
func NewNodeAlreadyNeighborError(node *Node) error {
	return &NodeAlreadyNeighborError{nodeName: node.Name()}
}

// Error - Implements error interface
func (e *NodeAlreadyNeighborError) Error() string {
	return fmt.Sprintf("Node %s already present in the newtowrk", e.nodeName)
}
