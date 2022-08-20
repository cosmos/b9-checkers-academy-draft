package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PlayerInfoKeyPrefix is the prefix to retrieve all PlayerInfo
	PlayerInfoKeyPrefix = "PlayerInfo/value/"
)

// PlayerInfoKey returns the store key to retrieve a PlayerInfo from the index fields
func PlayerInfoKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
