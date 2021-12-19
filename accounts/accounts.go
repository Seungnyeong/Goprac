package accounts

import (
	"errors"
	"fmt"
)

//Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of yout Account
func (a Account) Balance() int {
	return a.balance
}

// WIthdraw x amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

// 내부적으로 호출하는 방법
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas : ", a.Balance())
}
