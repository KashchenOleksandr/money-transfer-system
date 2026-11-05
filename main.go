package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be a positive value")
	}
	u.Balance += amount
	return nil
}
func (u *User) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > u.Balance {
		return errors.New("insufficient funds")
	}
	u.Balance -= amount
	return nil
}

func main() {
	user1 := &User{
		ID:      "1",
		Name:    "Megan",
		Balance: 100.0,
	}
	user2 := &User{
		ID:      "2",
		Name:    "Fox",
		Balance: 200.0,
	}
	if err := user1.Deposit(80.0); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user1.Name, user1.Balance)

	if err := user2.Withdraw(100.0); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user2.Name, user2.Balance)

	if err := user1.Withdraw(300.0); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user1.Name, user1.Balance)

	if err := user1.Withdraw(50.0); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user1.Name, user1.Balance)

	if err := user2.Deposit(50.0); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user2.Name, user2.Balance)
}
