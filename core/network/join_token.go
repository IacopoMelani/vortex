package network

import (
	"sync"
	"time"

	"github.com/IacopoMelani/vortex/utils"
	"github.com/google/uuid"
)

// MARK: consts

const (
	DefaultExpJoinToken = 60 * 5 // 5 minutes
)

// MARK: JoinToken, JoinTokenConfig & constructors

// JoinToken - Defines a struct for node join token
type JoinToken struct {
	sync.RWMutex
	id    string
	host  string
	value string
	iat   time.Time
	exp   time.Time
}

// JoinTokenConfig - Defines the JoinToken config for constructor
type JoinTokenConfig struct {
	Host string
}

// NewJoinToken - Returns a new join token
func NewJoinToken() (*JoinToken, error) {

	value, err := utils.SecureRandomString(64)
	if err != nil {
		return nil, err
	}

	ipAddr, err := utils.GetPrimaryIP()
	if err != nil {
		return nil, err
	}

	return &JoinToken{
		id:    uuid.New().String(),
		iat:   time.Now().UTC(),
		exp:   time.Now().UTC().Add(DefaultExpJoinToken * time.Second),
		host:  ipAddr.String(),
		value: value,
	}, nil
}

// NewJoinTokenWithConfig - Returns a new JoinToken instance with config param, see JoinTokenConfig
func NewJoinTokenWithConfig(jtConfig JoinTokenConfig) (*JoinToken, error) {

	jt, err := NewJoinToken()
	if err != nil {
		return nil, err
	}

	if jtConfig.Host != "" {
		jt.host = jtConfig.Host
	}

	return jt, nil
}

// MARK: JoinToken exported

func (j *JoinToken) Host() string {
	j.RLock()
	defer j.RUnlock()
	return j.host
}

// ID - Returns JoinToken ID
func (j *JoinToken) ID() string {
	j.RLock()
	defer j.RUnlock()
	return j.id
}

// Value - Returns JoinToken Value
func (j *JoinToken) Value() string {
	j.RLock()
	defer j.RUnlock()
	return j.value
}
