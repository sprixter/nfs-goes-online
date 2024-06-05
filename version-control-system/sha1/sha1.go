package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

// hashSHA1 takes an input string and returns its SHA1 hash value as a hex-encoded string.
func hashSHA1(input string) string {
	// Create a new SHA1 hasher.
	hasher := sha1.New()

	// Write the input data to the hasher.
	hasher.Write([]byte(input))

	// Compute the SHA1 hash.
	hashBytes := hasher.Sum(nil)

	// Encode the hash bytes to a hex string.
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
