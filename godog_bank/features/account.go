// account.go

package bank

import "fmt"

type account struct {
	balance int
}

func (a *account) withdraw(amount int) error {
	if a.balance < amount {
		return fmt.Errorf("Insufficient Funds! Available Balance: %v$", a.balance)
	}

	a.balance -= amount
	return nil
}

func (a *account) deposit(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("Not a valid amount to deposit: %v$", amount)
	}

	a.balance += amount
	return nil
}

func (a *account) transfer(destination *account, amount int) error {

	if a.balance >= amount {
		destination.balance += amount
		a.withdraw(amount)
	} else {
		return fmt.Errorf("Insufficient funds! Requested Amount: %v$, Available Balance: %v$", amount, a.balance)
	}
	return nil
}
