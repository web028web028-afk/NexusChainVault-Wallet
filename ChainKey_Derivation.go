package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func DeriveSecp256k1Key() (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func PublicKeyToAddress(pub *ecdsa.PublicKey) string {
	pubBytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	addrBytes := pubBytes[1:]
	return "0x" + hexEncode(addrBytes)[24:]
}

func hexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

func ValidatePrivateKeyFormat(key *ecdsa.PrivateKey) bool {
	return key.D != nil && key.D.Cmp(big.NewInt(0)) > 0
}
