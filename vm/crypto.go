package vm

import (
	"crypto/sha1"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Sha1 computes the sha1 hash of the given bytes.
func Sha1(b []byte) []byte {
	sha := sha1.New()
	sha.Write(b)
	return sha.Sum(nil)
}

// Sha256 computes the sha256 hash of the given bytes.
func Sha256(b []byte) []byte {
	sha := sha256.New()
	sha.Write(b)
	return sha.Sum(nil)
}

// Hash256 computes the 2^sha256 hash of the given bytes.
func Hash256(b []byte) []byte {
	sha := sha256.New()
	sha.Write(b)
	hash := sha.Sum(nil)
	sha.Reset()
	sha.Write(hash)
	return sha.Sum(nil)
}

// Hash160 computes the RIPEMD160 hash over the sha256 hash of the given bytes.
func Hash160(b []byte) []byte {
	hash := Sha256(b)
	ripe := ripemd160.New()
	ripe.Write(hash)
	return ripe.Sum(nil)
}
