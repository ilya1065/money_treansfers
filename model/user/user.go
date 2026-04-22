package user

import (
	"errors"
	"sync"
)

type User struct {
	mu      sync.Mutex
	ID      string
	Name    string
	Balance float64
}

func NewUser(id, name string, balance float64) *User {
	return &User{
		ID:      id,
		Name:    name,
		Balance: balance,
	}
}

func (u *User) Deposit(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.Balance < amount {
		return errors.New("Недостаточно средств на счету")
	}
	u.Balance -= amount
	return nil
}
