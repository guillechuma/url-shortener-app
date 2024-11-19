package url

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(originalURL string) string {
	// shorten the URL. Use hashing, deterministic.
	h := sha256.New()
	h.Write([]byte(originalURL))
	hash := hex.EncodeToString(h.Sum(nil))
	// Shorten the hash to first 8 characters
	shortURL := hash[:8]
	return shortURL
}
