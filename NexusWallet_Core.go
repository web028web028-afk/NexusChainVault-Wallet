package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

type NexusWallet struct {
	WalletID     string
	ChainType    string
	EncryptedKey string
	Address      string
}

func GenerateWalletID() string {
	hash := sha256.New()
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)
	hash.Write(randomBytes)
	return hex.EncodeToString(hash.Sum(nil))[:16]
}

func EncryptKey(plainKey, password string) (string, error) {
	block, err := aes.NewCipher([]byte(generateAESKey(password)))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainKey), nil)
	return hex.EncodeToString(ciphertext), nil
}

func generateAESKey(pwd string) []byte {
	hash := sha256.Sum256([]byte(pwd))
	return hash[:]
}

func (w *NexusWallet) ValidateAddress() bool {
	return len(w.Address) == 42 && w.Address[:2] == "0x"
}

func main() {
	wallet := NexusWallet{
		WalletID:  GenerateWalletID(),
		ChainType: "EVM",
	}
}
