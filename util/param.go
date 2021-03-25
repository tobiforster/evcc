package util

import (
	"strconv"

	"github.com/andig/evcc/core/msg"
)

// Param is the broadcast channel data type
type Param struct {
	LoadPoint *int
	Key       msg.Message
	Val       interface{}
}

// UniqueID returns unique identifier for parameter LoadPoint/Key combination
func (p Param) UniqueID() string {
	key := p.Key.Key()
	if p.LoadPoint != nil {
		key = strconv.Itoa(*p.LoadPoint) + "." + key
	}
	return key
}
