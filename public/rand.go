package main

import (
	"math/rand"
	"time"
)

// Possible characters that the random string may contain
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Creates source for randomization based on the Unix timestamp
var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// Returns random int between ints min and max
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Returns random string with specified length and using a character set
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Returns random string with specified length, and uses variable charset instead of a specified character set
func randString(length int) string {
	return stringWithCharset(length, charset)
}
