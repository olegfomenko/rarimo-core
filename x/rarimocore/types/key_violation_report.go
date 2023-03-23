package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// ViolationReportKeyPrefix is the prefix to retrieve all ViolationReport
	ViolationReportKeyPrefix = "ViolationReport/value/"
)

// ViolationReportKey returns the store key to retrieve a ViolationReport from the index fields
func ViolationReportKey(
	index *ViolationReportIndex,
) []byte {
	var key []byte

	key = append(key, []byte(index.SessionId)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Offender)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(fmt.Sprintf("%d", index.ViolationType))...)
	key = append(key, []byte("/")...)

	return key
}
