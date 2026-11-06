package main

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	ID      string
	Name    string
	Balance float64
	mu      sync.Mutex
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users            map[string]*User
	TransactionQueue []Transaction
}

func (p *PaymentSystem) AddUser(user *User) {
	p.Users[user.ID] = user
}

func (p *PaymentSystem) AddTransaction(transaction Transaction) {
	p.TransactionQueue = append(p.TransactionQueue, transaction)
}

func (u *User) Deposit(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if amount <= 0 {
		return errors.New("deposit amount must be a positive value")
	}
	u.Balance += amount
	return nil

}

func (u *User) Withdraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > u.Balance {
		return errors.New("insufficient funds")
	}
	u.Balance -= amount
	return nil
}

func (p *PaymentSystem) ProcessingTransactions(transaction Transaction) error {
	if transaction.Amount <= 0 {
		return errors.New("transaction amount must be positive")
	}
	sender, ok := p.Users[transaction.FromID]
	if !ok {
		return errors.New("sender not found")
	}
	receiver, ok := p.Users[transaction.ToID]
	if !ok {
		return errors.New("receiver not found")
	}
	if err := sender.Withdraw(transaction.Amount); err != nil {
		return err
	}
	if err := receiver.Deposit(transaction.Amount); err != nil {
		sender.Deposit(transaction.Amount)
		return err
	}
	return nil
}

func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range ch {
		if err := ps.ProcessingTransactions(t); err != nil {
			fmt.Println("Error processing transaction:", err)
		}
	}
}

func main() {

	system := &PaymentSystem{
		Users:            make(map[string]*User),
		TransactionQueue: []Transaction{},
	}

	user1 := &User{
		ID:      "1",
		Name:    "Megan",
		Balance: 100.0,
	}
	system.AddUser(user1)

	user2 := &User{
		ID:      "2",
		Name:    "Fox",
		Balance: 200.0,
	}
	system.AddUser(user2)

	transaction1 := Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 50.0,
	}
	system.AddTransaction(transaction1)

	transaction2 := Transaction{
		FromID: "2",
		ToID:   "1",
		Amount: 30.0,
	}
	system.AddTransaction(transaction2)

	transaction3 := Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 222.0,
	}

	system.AddTransaction(transaction3)

	wg := &sync.WaitGroup{}

	var ch chan Transaction = make(chan Transaction, len(system.TransactionQueue))

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go system.Worker(ch, wg)
	}

	for _, t := range system.TransactionQueue {
		ch <- t
	}

	close(ch)

	wg.Wait()

	fmt.Println(user1.Name, user1.Balance)
	fmt.Println(user2.Name, user2.Balance)
}
