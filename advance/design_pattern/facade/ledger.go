package main

import "fmt"

type ledger struct{}

func (s *ledger) makeEntry(accountID, txtType string, amount int) {
	fmt.Printf("Make ledger entry for accoutId %s with txtType %s for amount %d", accountID, txtType, amount)
	return
}
