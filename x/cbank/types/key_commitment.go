package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// CommitmentKeyPrefix is the prefix to retrieve all Commitment
	CommitmentKeyPrefix = "Commitment/value/"
)

// CommitmentKey returns the store key to retrieve an Commitment from the index fields
func CommitmentKey(
	index string,
) []byte {
	var key []byte

	key = append(key, []byte(index)...)
	key = append(key, []byte("/")...)

	return key
}
