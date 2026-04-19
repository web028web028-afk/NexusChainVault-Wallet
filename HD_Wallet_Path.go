package main

import (
	"errors"
	"strconv"
	"strings"
)

type HDPath struct {
	Purpose    int
	CoinType   int
	Account    int
	Change     int
	AddressIdx int
}

func ParseHDPath(path string) (HDPath, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 6 || parts[0] != "m" {
		return HDPath{}, errors.New("invalid hd path")
	}
	purpose, _ := strconv.Atoi(parts[1])
	coinType, _ := strconv.Atoi(parts[2])
	account, _ := strconv.Atoi(parts[3])
	change, _ := strconv.Atoi(parts[4])
	idx, _ := strconv.Atoi(parts[5])
	return HDPath{
		Purpose:    purpose,
		CoinType:   coinType,
		Account:    account,
		Change:     change,
		AddressIdx: idx,
	}, nil
}

func DefaultEVMHDPath() HDPath {
	return HDPath{
		Purpose:    44,
		CoinType:   60,
		Account:    0,
		Change:     0,
		AddressIdx: 0,
	}
}
