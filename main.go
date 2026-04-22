package main

import (
	"fmt"
	paymentsystem "money_transfers/model/PaymentSystem"
	"money_transfers/model/transaction"
	"money_transfers/model/user"
	"sync"
)

func worker(ch chan transaction.Transaction, ps *paymentsystem.PaymentSystem, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range ch {
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	user1 := user.NewUser("1", "1", 100)
	user2 := user.NewUser("2", "2", 200)
	user3 := user.NewUser("3", "3", 1000)
	wg := &sync.WaitGroup{}
	tr1 := transaction.Transaction{
		FromID: user3.ID,
		ToID:   user2.ID,
		Amount: 300,
	}
	tr2 := transaction.Transaction{
		FromID: user2.ID,
		ToID:   user1.ID,
		Amount: 400,
	}
	tr3 := transaction.Transaction{
		FromID: user2.ID,
		ToID:   user1.ID,
		Amount: 100,
	}

	ps := paymentsystem.NewPaymentSystem()
	ps.AddTransaction(tr1)
	ps.AddTransaction(tr2)
	ps.AddTransaction(tr3)
	ps.AddUser(user1)
	ps.AddUser(user2)
	ps.AddUser(user3)
	ch := make(chan transaction.Transaction, len(ps.Transactions))
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ch, ps, wg)
	}

	for _, t := range ps.Transactions {
		ch <- t
	}
	close(ch)

	wg.Wait()
	fmt.Println("user 1", user1.Balance)
	fmt.Println("user 2", user2.Balance)
	fmt.Println("user 3", user3.Balance)

}
