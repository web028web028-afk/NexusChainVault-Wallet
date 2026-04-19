package main

import (
	"archive/zip"
	"os"
)

type BackupManager struct {
	Wallet *NexusWallet
}

func (b *BackupManager) ExportEncryptedBackup(outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	f, _ := zipWriter.Create("wallet_backup.nexus")
	backupData := []byte(b.Wallet.EncryptedKey + "|" + b.Wallet.Address)
	_, _ = f.Write(backupData)
	return nil
}

func (b *BackupManager) ImportBackup(filePath string) error {
	return nil
}
