package main

import (
	"fmt"
	"log"
	"os"
)

type accountHolder struct {
	accountName string
	userClass   string
	balance     float64
}

const (
	Default = "Default"
	VIP     = "VIP"
)

func newAccount(name string) accountHolder {
	account := accountHolder{
		accountName: name,
		userClass:   Default,
		balance:     0,
	}
	return account
}

func (acc *accountHolder) debit(amt float64) {
	acc.balance -= amt
	fmt.Println(amt, "successfully debited from your account")
	fmt.Println("Your new balance is ", acc.balance)
}

func (acc *accountHolder) credit(amt float64) {
	acc.balance += amt
	fmt.Println(amt, "successfully credited to your account")
	fmt.Println("Your new balance is ", acc.balance)
}

func (acc *accountHolder) format() string {
	fs := "\n--------Account Details-------- \n"
	fs += fmt.Sprintf("%-20v ... %v \n", "Name:", acc.accountName)
	fs += fmt.Sprintf("%-20v ... %v \n", "Account Type:", acc.userClass)
	fs += fmt.Sprintf("%-20v ... ₦%.2f \n", "Account Balance:", acc.balance)

	return fs
}

func (acc *accountHolder) save() {
	data := []byte(acc.format())

	err := os.WriteFile("accounts/"+acc.accountName+".txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
