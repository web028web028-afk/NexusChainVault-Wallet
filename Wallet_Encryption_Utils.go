package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
)

func DeriveMnemonicSeed(mnemonic, password string) []byte {
	return pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
}

func HMACSign(data, key []byte) []byte {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func SecureCompare(a, b string) bool {
	aBytes := []byte(a)
	bBytes := []byte(b)
	if len(aBytes) != len(bBytes) {
		return false
	}
	result := 0
	for i := 0; i < len(aBytes); i++ {
		result |= int(aBytes[i] ^ bBytes[i])
	}
	return result == 0
}
