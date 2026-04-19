package main

import "fmt"

type Notification struct {
	Title     string
	Content   string
	Timestamp int64
	Type      string
}

type NotifyService struct {
	Notifications []Notification
}

func (n *NotifyService) PushTxSuccess(txHash string) {
	n.Notifications = append(n.Notifications, Notification{
		Title:     "Transaction Confirmed",
		Content:   fmt.Sprintf("Tx %s confirmed successfully", txHash),
		Timestamp: time.Now().Unix(),
		Type:      "success",
	})
}

func (n *NotifyService) PushTxFailed(txHash string) {
	n.Notifications = append(n.Notifications, Notification{
		Title:     "Transaction Failed",
		Content:   fmt.Sprintf("Tx %s execution failed", txHash),
		Timestamp: time.Now().Unix(),
		Type:      "error",
	})
}
