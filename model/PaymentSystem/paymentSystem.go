package paymentsystem

import (
	"errors"
	"money_transfers/model/transaction"
	"money_transfers/model/user"
)

type PaymentSystem struct {
	Users        map[string]*user.User
	Transactions []transaction.Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:        make(map[string]*user.User),
		Transactions: []transaction.Transaction{},
	}
}

func (ps *PaymentSystem) AddUser(user *user.User) {
	ps.Users[user.ID] = user

}

func (ps *PaymentSystem) AddTransaction(transaction transaction.Transaction) {
	ps.Transactions = append(ps.Transactions, transaction)
}

func (ps *PaymentSystem) ProcessingTransactions(transaction transaction.Transaction) error {
	from, ok := ps.Users[transaction.FromID]
	if !ok {
		return errors.New("пользователь не найден")
	}
	to, ok := ps.Users[transaction.ToID]
	if !ok {
		return errors.New("пользователь не найден")
	}
	err := from.Withdraw(transaction.Amount)
	if err != nil {
		return err
	}
	to.Deposit(transaction.Amount)
	return nil
}
