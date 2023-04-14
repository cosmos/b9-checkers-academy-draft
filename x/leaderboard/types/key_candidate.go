package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	CandidateKeyPrefix = "Candidate/value/"
)

// CandidateKey returns the store key to retrieve a Candidate from the index field
// It is not used but is here to remind where values are stored
func CandidateKey(address []byte) []byte {
	var key []byte

	prefixBytes := []byte(CandidateKeyPrefix)
	key = append(key, prefixBytes...)
	key = append(key, address...)

	return key
}
