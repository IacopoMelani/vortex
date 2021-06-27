package network

import (
	"fmt"
	"os"
	"sync"

	"github.com/IacopoMelani/vortex-storage/utils"
	"github.com/google/uuid"
)

// MARK: consts

const (
	DefaultRPCPort = ":6414"
)

// MARK: Node, NodeConfig & constructors

// Node - Defines a node of vortex network
type Node struct {
	sync.RWMutex
	neighbors  map[string]*Node
	joinTokens map[string]*JoinToken
	host       string
	id         string
	name       string
	rpcPort    string
}

// NodeConfig - Defines a node config struct
type NodeConfig struct {
	Name    string
	IP      string
	RPCPort string
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
		id:         uuid.New().String(),
		name:       hostname,
		host:       ip.String(),
		rpcPort:    DefaultRPCPort,
		neighbors:  make(map[string]*Node),
		joinTokens: make(map[string]*JoinToken),
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
	if config.RPCPort != "" {
		node.rpcPort = config.RPCPort
	}
	if config.IP != "" {
		node.host = config.IP
	}

	return node, nil
}

// MARK: Node exported

// AddNeighbor - Add new node in the neighbors networks
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

// JoinToken - Return a new JoinToken
func (n *Node) JoinToken() (*JoinToken, error) {

	jt, err := NewJoinToken()
	if err != nil {
		return nil, err
	}

	n.Lock()
	defer n.Unlock()

	n.joinTokens[jt.ID()] = jt

	return jt, nil
}

// MARK: NodeAlreadyNeighborError

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
