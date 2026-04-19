package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var wordList = []string{
	"apple", "banana", "cherry", "date", "elder", "fig", "grape", "honey",
	"ice", "juice", "kiwi", "lemon", "mango", "nut", "orange", "pear",
}

func Generate12WordMnemonic() (string, error) {
	words := make([]string, 12)
	for i := 0; i < 12; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))
		if err != nil {
			return "", err
		}
		words[i] = wordList[idx.Int64()]
	}
	return strings.Join(words, " "), nil
}

func ValidateMnemonic(phrase string) bool {
	words := strings.Fields(phrase)
	return len(words) == 12 || len(words) == 24
}
