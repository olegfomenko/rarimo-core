package operations

import (
	"bytes"

	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

// TransferOperation defines the token transfer operation - from one network to another with bundle call.
// Use for EVM networks where token name, symbol and decimals are already defined by parent contract
type TransferOperation struct {
	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
	// Target metadata information
	TargetURI string
	Bundle    []byte
	Salt      []byte
}

func NewTransferOperation(addHex, idHex, amount, uri string) *TransferOperation {
	return &TransferOperation{
		TargetAddress: tryHexDecode(addHex),
		TargetId:      tryHexDecode(idHex),
		Amount:        amountBytes(amount),
		TargetURI:     uri,
		Bundle:        []byte{},
		Salt:          []byte{},
	}
}

func NewTransferWithBundleOperation(addHex, idHex, amount, uri, bundleData, bundleSalt string) *TransferOperation {
	return &TransferOperation{
		TargetAddress: tryHexDecode(addHex),
		TargetId:      tryHexDecode(idHex),
		Amount:        amountBytes(amount),
		TargetURI:     uri,
		Bundle:        tryHexDecode(bundleData),
		Salt:          tryHexDecode(bundleSalt),
	}
}

var _ Operation = &TransferOperation{}

func (t TransferOperation) GetContent() crypto.ContentData {
	return bytes.Join([][]byte{t.TargetAddress, t.TargetId, t.Amount, []byte(t.TargetURI), t.Bundle, t.Salt}, []byte{})
}

// TransferFullMetaOperation defines the token transfer operation - from one network to another with full token metadata
type TransferFullMetaOperation struct {
	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount []byte
	// Target metadata information
	TargetName     string
	TargetSymbol   string
	TargetURI      string
	TargetDecimals []byte
	Bundle         []byte
	Salt           []byte
}

func NewTransferFullMetaOperation(addHex, idHex, amount, name, symbol, uri string, decimals uint8) *TransferFullMetaOperation {
	return &TransferFullMetaOperation{
		TargetAddress:  tryHexDecode(addHex),
		TargetId:       tryHexDecode(idHex),
		Amount:         amountBytes(amount),
		TargetName:     name,
		TargetSymbol:   symbol,
		TargetURI:      uri,
		TargetDecimals: decimalsBytes(decimals),
		Bundle:         []byte{},
		Salt:           []byte{},
	}
}

func NewTransferFullMetaWithBundleOperation(addHex, idHex, amount, name, symbol, uri string, decimals uint8, bundleData, bundleSalt string) *TransferFullMetaOperation {
	return &TransferFullMetaOperation{
		TargetAddress:  tryHexDecode(addHex),
		TargetId:       tryHexDecode(idHex),
		Amount:         amountBytes(amount),
		TargetName:     name,
		TargetSymbol:   symbol,
		TargetURI:      uri,
		TargetDecimals: decimalsBytes(decimals),
		Bundle:         tryHexDecode(bundleData),
		Salt:           tryHexDecode(bundleSalt),
	}
}

var _ Operation = &TransferFullMetaOperation{}

func (t TransferFullMetaOperation) GetContent() crypto.ContentData {
	return bytes.Join([][]byte{t.TargetAddress, t.TargetId, t.Amount, []byte(t.TargetName), []byte(t.TargetSymbol), []byte(t.TargetURI), t.TargetDecimals, t.Bundle, t.Salt}, []byte{})
}
