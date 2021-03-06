package pwdhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"math"
)

const (
	MinCost int = 1             // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost int = math.MaxInt32 // the maximum allowable cost as passed in to GenerateFromPassword
)

const (
	base    int    = 10
	bitSize int    = 64
	null    string = "\x00"

	algorithmIndex int = 0
	costIndex      int = 1
	saltIndex      int = 2
	digestIndex    int = 3
)

var (
	delim []byte = []byte{'$'}

	// supported hash functions
	algorithms = map[string]func() hash.Hash{
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}
)
