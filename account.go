package main

//Import requisite libraries
import (
	"fmt"
	"log"
	"os"
)

//declare constant values for userclass
const (
	Default = "Default"
	VIP     = "VIP"
)

//declare accountHolder structure
type accountHolder struct {
	accountName string
	userClass   string
	balance     float64
	PIN         int
}

//function to make new accountHolder types
func newAccount(name string) accountHolder {
	account := accountHolder{
		accountName: name,
		userClass:   Default,
		balance:     0,
	}
	return account
}

//receiver function on accountholder to debit balance
func (acc *accountHolder) debit(amt float64) {
	acc.balance -= amt
	fmt.Println(amt, "successfully debited from your account")
	fmt.Println("Your new balance is ", acc.balance)
}

//receiver function on accountholder to credit balance
func (acc *accountHolder) credit(amt float64) {
	acc.balance += amt
	fmt.Println(amt, "successfully credited to your account")
	fmt.Println("Your new balance is ", acc.balance)
}

//receiver function on accountHolder information
func (acc *accountHolder) format() string {
	fs := ""
	switch acc.accountName {
	case "":
		fs = "\nNo Account to Format\n"
	default:
		fs = "\n--------Account Details-------- \n"
		fs += fmt.Sprintf("%-20v ... %v \n", "Name:", acc.accountName)
		fs += fmt.Sprintf("%-20v ... %v \n", "Account Type:", acc.userClass)
		fs += fmt.Sprintf("%-20v ... ₦%.2f \n", "Account Balance:", acc.balance)
	}
	return fs
}

//receiver function on accountholder to save accountHolder information to txt file
func (acc *accountHolder) save() {
	data := []byte(acc.format())

	err := os.WriteFile("accounts/"+acc.accountName+".txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}