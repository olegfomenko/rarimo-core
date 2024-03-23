package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// PendingTransferKeyPrefix is the prefix to retrieve all PendingTransfer
	PendingTransferKeyPrefix = "PendingTransfer/value/"
)

// PendingTransferKey returns the store key to retrieve an PendingTransfer from the index fields
func PendingTransferKey(
	pt *PendingTransfer,
) []byte {
	var key []byte

	key = append(key, []byte(pt.Index)...)
	key = append(key, []byte("/")...)

	return key
}
