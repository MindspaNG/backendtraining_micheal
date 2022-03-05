package main

// import requisite libraries
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//function to get input from console
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

//function to create new accountHolders
func createAccountHolder(user map[int64]string) accountHolder {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nChoose an Account Type (0 - default, 1 - VIP, 99 - Cancel): ", reader)
	var acc accountHolder

	switch opt {
	case "0":
		fmt.Println("Default Account selected")
		amt, _ := getInput("\nDeposit a minimum of NGN 10,000 to open your account: ", reader)

		a, err := strconv.ParseFloat(amt, 64)
		if err != nil {
			fmt.Println("Deposit amount must be a number!")
			createAccountHolder(user)
		} else if a < 10000.00 {
			fmt.Println("You must deposit at least NGN 10,000")
			createAccountHolder(user)
		} else {
			userName, _ := getInput("Enter Account Name:", reader)
			acc = newAccount(userName)
			acc.balance += a
			fmt.Println("\nAccount Created Successfully!")
			fmt.Println(acc.format())
			acc.save()
		}

	case "1":
		fmt.Println("VIP Account selected")
		userName := accountValidation(user)
		switch userName {
		case "":
			acc.accountName = ""
		default:
			acc.accountName = userName
			acc.balance += 50000.00
			acc.userClass = VIP
			fmt.Println("\n Account Created Successfully!")
			fmt.Println(acc.format())
			acc.save()

		case "99":
			fmt.Println("Thank you!!")
		}
	}
	return acc
}

// function to validate VIP user before account creation
func accountValidation(user map[int64]string) string {
	reader := bufio.NewReader(os.Stdin)
	pin, _ := getInput("\nEnter PIN: ", reader)
	p, err := strconv.ParseInt(pin, 10, 64)
	p1, found := user[p]
	vip := ""
	if err != nil {
		fmt.Println("PIN must be numeric!")
		accountValidation(user)
	} else {
		switch found {
		case true:
			fmt.Println("User", p1, "found")
			vip = p1
		case false:
			fmt.Println("PIN does not exist!")
			bl := YesNo("\nTry Again?")
			switch bl {
			case true:
				accountValidation(user)
			case false:
				createAccountHolder(user)
			}
		}
	}
	return vip
}

// function to deposit into account
func deposit(acc *accountHolder) {
	switch acc.accountName {
	case "":
		fmt.Println("\nYou cannot deposit, You Do not have an account Yet")
	default:
		reader := bufio.NewReader(os.Stdin)
		bl := YesNo("Do You want to make a deposit? ")
		switch bl {
		case true:
			amt, _ := getInput("\nEnter Amount to Deposit: ", reader)
			a, err := strconv.ParseFloat(amt, 64)
			if err != nil {
				fmt.Println("Deposit amount must be a number!")
				deposit(acc)
			} else {
				acc.credit(a)
				acc.save()
			}
		case false:
			fmt.Println("Thank You")
		}
	}
}

// function to withdraw from account
func withdraw(acc *accountHolder) {
	switch acc.accountName {
	case "":
		fmt.Println("\nYou caanot make a withdrawal, You Do not have an account Yet")
	default:
		reader := bufio.NewReader(os.Stdin)
		bl := YesNo("Do You want to make a Withdrawal? ")
		switch bl {
		case true:
			amt, _ := getInput("\nEnter Amount to withdraw: ", reader)
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
		case false:
			fmt.Println("Thank You")
		}
	}
}

//function to get Yes or No input from console
func YesNo(msg string) bool {
	reader := bufio.NewReader(os.Stdin)
	sel, _ := getInput(msg+"Y - Yes,  N - No  ", reader)
	s := strings.ToUpper(sel)
	bl := true

	switch s {
	case "Y":
		bl = true
	case "N":
		bl = false
	default:
		fmt.Println("Select the right option!")
	}
	return bl

}
func main() {
	// Pre-defined VIP users with PIN and Names
	PIN := map[int64]string{
		159: "Tim Jones",
		753: "Sam Smith",
		258: "Ada Stone",
		953: "Ama Young",
		751: "Kofi Poku",
	}

	// Create new accounts, verify PIN before creating VIP user
	myAccount := createAccountHolder(PIN)

	//Print account details to console
	fmt.Println(myAccount.format())

	//withdraw from account
	withdraw(&myAccount)

	//deposit into account
	deposit(&myAccount)

	//Print account details into console
	fmt.Println(myAccount.format())

}
