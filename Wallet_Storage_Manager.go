package main

import (
	"encoding/json"
	"os"
)

type WalletStorage struct {
	Wallets []NexusWallet `json:"wallets"`
	Path    string        `json:"-"`
}

func NewWalletStorage(path string) *WalletStorage {
	return &WalletStorage{Path: path}
}

func (ws *WalletStorage) Save() error {
	data, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ws.Path, data, 0600)
}

func (ws *WalletStorage) Load() error {
	data, err := os.ReadFile(ws.Path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, ws)
}

func (ws *WalletStorage) AddWallet(w NexusWallet) {
	ws.Wallets = append(ws.Wallets, w)
}
