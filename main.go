package main

import (
	"fmt"
	"log"

	"github.com/snkim/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
