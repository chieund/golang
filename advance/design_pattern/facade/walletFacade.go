package main

import (
	"facade/folder"
	"fmt"
)

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notificaltion
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notificaltion{},
		ledger:       &ledger{},
	}
	fmt.Println("Account Created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creaditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit monney from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNOtification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
