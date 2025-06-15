package id

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GenInstanceID(hashName string, uid uint64, prefix string, len int) string {
	hash := sha256.New()
	hash.Write([]byte(hashName + strconv.Itoa(int(uid))))
	hashStr := hex.EncodeToString(hash.Sum(nil))

	return prefix + hashStr[:len]
}

type Alphabet string

// Defiens alphabet.
const (
	Alphabet62 Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func RandString(letters Alphabet, n int) string {
	buf := make([]byte, n)
	randomness := make([]byte, n)

	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(letters)
	for pos := range buf {
		randon := randomness[pos]

		randomPos := randon % uint8(l)

		buf[pos] = letters[randomPos]
	}

	return string(buf)
}

func NewRandonStr62(n int) string {
	return RandString(Alphabet62, n)
}

func NewRandonStr36(n int) string {
	return RandString(Alphabet36, n)
}
