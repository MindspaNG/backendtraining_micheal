package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PIN uint64 = 1235

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createAccountHolder() accountHolder {
	reader := bufio.NewReader(os.Stdin)

	userName, _ := getInput("Create a new Account:", reader)

	acc := newAccount(userName)
	fmt.Println("User created - ", acc.accountName)

	return acc
}

func accountOptions(acc *accountHolder) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nChoose an Account Type (0 - default, 1 - VIP): ", reader)

	switch opt {
	case "0":
		fmt.Println("Default Account selected")
		amt, _ := getInput("\nDeposit a minimum of NGN 10,000 to open your account: ", reader)

		a, err := strconv.ParseFloat(amt, 64)
		if err != nil {
			fmt.Println("Deposit amount must be a number!")
			accountOptions(acc)
		} else if a < 10000.00 {
			fmt.Println("You must deposit at least NGN 10,000")
			accountOptions(acc)
		} else {
			acc.balance += a
			fmt.Println("\n Account Created Successfully!")
			fmt.Println(acc.format())
			acc.save()
		}

	case "1":
		fmt.Println("VIP Account selected")
		pin, _ := getInput("\nEnter your PIN to validate your VIP account: ", reader)

		p, err := strconv.ParseUint(pin, 10, 64)
		if err != nil {
			fmt.Println("Enter a valid PIN")
			accountOptions(acc)
		} else if p != PIN {
			fmt.Println("You have entered a wrong Pin. Please try again ")
			accountOptions(acc)
		} else if p == PIN {
			fmt.Println("Your account validation was successful")
			acc.balance += 50000.00
			acc.userClass = VIP
			fmt.Println("\nAccount Created Successfully!")
			fmt.Println(acc.format())
			acc.save()
		}

	default:
		fmt.Println("Select the correct option")
		accountOptions(acc)
	}
}

func deposit(acc *accountHolder) {
	reader := bufio.NewReader(os.Stdin)
	amt, _ := getInput("\nEnter Amount to Deposit: ", reader)
	a, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		fmt.Println("Deposit amount must be a number!")
		deposit(acc)
	} else {
		acc.credit(a)
		acc.save()
	}
}

func withdraw(acc *accountHolder) {
	reader := bufio.NewReader(os.Stdin)
	amt, _ := getInput("\nEnter Amount to Withdraw: ", reader)
	a, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		fmt.Println("Withdrawal amount must be a number!")
		withdraw(acc)
	} else if a > acc.balance {
		fmt.Println("Your balance is not enough to withdraw! Balance: ", acc.balance)
		withdraw(acc)
	} else {
		acc.debit(a)
		acc.save()
	}
}
func main() {
	myAccount := createAccountHolder()
	accountOptions(&myAccount)
	fmt.Println(myAccount.format())
	withdraw(&myAccount)
	fmt.Println(myAccount.format())
	deposit(&myAccount)
	fmt.Println(myAccount.format())

}
