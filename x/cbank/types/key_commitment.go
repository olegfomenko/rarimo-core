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
	com *Commitment,
) []byte {
	var key []byte

	key = append(key, com.Point.CompressBytes()...)
	key = append(key, []byte("/")...)

	return key
}
