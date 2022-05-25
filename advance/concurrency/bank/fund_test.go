package main

import (
	"fmt"
	"testing"
)

func BenchmarkFund(b *testing.B) {
	fund := NewFund(b.N)
	fmt.Println("N=", b.N)
	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	println(fund.Balance())

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}
