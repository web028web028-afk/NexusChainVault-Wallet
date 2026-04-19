package main

import (
	"sync"
	"time"
)

type WalletLock struct {
	isLocked bool
	mutex    sync.Mutex
	lockTime time.Time
	timeout  time.Duration
}

func NewWalletLock(timeoutMin int) *WalletLock {
	return &WalletLock{
		isLocked: false,
		timeout:  time.Duration(timeoutMin) * time.Minute,
	}
}

func (wl *WalletLock) LockWallet() {
	wl.mutex.Lock()
	defer wl.mutex.Unlock()
	wl.isLocked = true
	wl.lockTime = time.Now()
}

func (wl *WalletLock) UnlockWallet(password string) bool {
	if SecureCompare(password, "wallet_secure_pass") {
		wl.isLocked = false
		return true
	}
	return false
}

func (wl *WalletLock) AutoLockCheck() {
	if time.Since(wl.lockTime) > wl.timeout {
		wl.LockWallet()
	}
}
