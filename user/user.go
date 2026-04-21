package user

import (
	"errors"
)


type User struct{
	ID string
	Name string
	Balance float64
}

func NewUser(id , name string, balance float64) User{
	return User{
		ID: id,
		Name:name,
		Balance: balance,
	}
}


func (u *User) Deposit (amount float64){
	u.Balance += amount
}

func (u *User) Withdraw (amount float64)error{
	if u.Balance < amount{
		return errors.New("Недостаточно средств на счету\n")
	}
	u.Balance -= amount
	return nil
}
