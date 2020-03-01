package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errorNoMoney = errors.New("can't withDraw")

// NewAccount can make a account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// WithDraw x amount from account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errorNoMoney
	}
	a.balance -= amount
	return nil
}

// Deposite x amount from account
func (a *Account) Deposite(amount int) {
	a.balance += amount
}

func (a *Account) String() {
	fmt.Println("name=", a.owner, ", balance=", a.balance)
}
