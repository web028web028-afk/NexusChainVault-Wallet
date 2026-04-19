package main

import (
	"crypto/ecdsa"
	"math/big"
)

func SignTransactionHash(hash []byte, priv *ecdsa.PrivateKey) (r, s *big.Int, v byte, err error) {
	pr, ps, pv, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		return nil, nil, 0, err
	}
	return pr, ps, byte(pv), nil
}

func VerifySignature(hash []byte, pub *ecdsa.PublicKey, r, s *big.Int) bool {
	return ecdsa.Verify(pub, hash, r, s)
}

func SerializeSignature(r, s *big.Int, v byte) []byte {
	sig := make([]byte, 65)
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = v
	return sig
}
