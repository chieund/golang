package main

import "fmt"

type notificaltion struct {
}

func (n *notificaltion) sendWalletCreditNotification() {
	fmt.Println("Send wallet credit notification")
}

func (n *notificaltion) sendWalletDebitNOtification() {
	fmt.Println("Sending wallet debit notification")
}
