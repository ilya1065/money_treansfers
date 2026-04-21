package main

import (
	"fmt"
	"money_transfers/user"
)

func main() {
	user1 := user.NewUser("1","1",100)
	user2 := user.NewUser("2","2",200)
	user3 := user.NewUser("3", "3",1000)

	fmt.Println("user1", user1.Balance)
	fmt.Println("user2", user2.Balance)
	fmt.Println("user3", user3.Balance)
	fmt.Println()


	user1.Deposit(1000)
	fmt.Println("user1 ",user1.Balance)
	user2.Deposit(10)
	fmt.Println("user2 ",user2.Balance)
	err := user3.Withdraw(1000)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("user3 ",user3.Balance)
	user3.Deposit(100)
	fmt.Println("user3 ",user3.Balance)
	err = user1.Withdraw(10000)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("user1 ",user1.Balance)
	err = user2.Withdraw(100)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("user2 ",user2.Balance)

	}
