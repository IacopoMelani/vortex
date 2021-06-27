package network

import (
	"sync"
	"time"

	"github.com/IacopoMelani/vortex-storage/utils"
	"github.com/google/uuid"
)

// MARK: consts

const (
	DefaultExpJoinToken = 60 * 5 // 5 minutes
)

// MARK: JoinToken & constructors

// JoinToken - Defines a struct for node join token
type JoinToken struct {
	sync.RWMutex
	id    string
	value string
	iat   time.Time
	exp   time.Time
}

// NewJoinToken - Returns a new join token
func NewJoinToken() (*JoinToken, error) {

	value, err := utils.SecureRandomString(64)
	if err != nil {
		return nil, err
	}

	return &JoinToken{
		id:    uuid.New().String(),
		iat:   time.Now().UTC(),
		exp:   time.Now().UTC().Add(DefaultExpJoinToken * time.Second),
		value: value,
	}, nil
}

// MARK: JoinToken exported

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
