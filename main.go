package main

import (
	"BankProject/accounts"
	"BankProject/costumers"
	"fmt"
	"os"
)

func Menu() {

	fmt.Println()
	fmt.Println("Hello !")
	fmt.Println("")

	fmt.Println("1 - Deposit")
	fmt.Println("2 - Withdraw")
	fmt.Println("3 - Transfer")
	fmt.Println("4 - Pay Bill")
	fmt.Println("5 - Exit")
}

func Options() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("")

	return option
}

func main() {
	caioProfile := costumers.Costumers{
		Name: "Caio",
		Age:  "21",
		CPF:  "042.078.411-05",
		Job:  "Developer"}
	caioAccount := accounts.BankData{
		Costumer:      caioProfile,
		AccountNumber: 1,
		AgencyNumber:  1,
		//Balance:       100.1
	}

	byaProfile := costumers.Costumers{
		Name: "Bya",
		Age:  "27",
		CPF:  "042.079.321-60",
		Job:  "Financial"}
	byaAccount := accounts.BankData{
		Costumer:      byaProfile,
		AccountNumber: 1,
		AgencyNumber:  1,
		//Balance:       50.
	}

	//

	//var cabralAccount *BankData
	//cabralAccount = new(BankData)
	//cabralAccount.client = "Cabral"
	//cabralAccount.balance = 123.32

	Menu()
	for {
		execMenu := Options()

		switch execMenu {
		case 1:
			fmt.Println("Type Deposit value:")
			var valueDeposit float64
			fmt.Scan(&valueDeposit)

			fmt.Println(caioAccount.Deposit(valueDeposit))

		case 2:

			fmt.Println("Type the withdraw amount:")
			var withdrawValue float64
			fmt.Scan(&withdrawValue)

			fmt.Println(caioAccount.Withdraw(withdrawValue))

		case 3:
			fmt.Println("Type the amount to transfer:")
			var transferValue float64

			fmt.Scan(&transferValue)

			fmt.Println(caioAccount.Transfer(transferValue, &byaAccount))

		case 4:

			fmt.Println("Type the bill cost:")
			var billValue float64

			fmt.Scan(&billValue)
			caioAccount.Deposit(100)
			caioAccount.PayBills(&caioAccount, billValue)
			fmt.Println(caioAccount.SeeBalance())
		case 5:
			fmt.Println("See you later")
			os.Exit(0)

		default:
			fmt.Printf("Unknown command! Choose a number between 1 and 5")
		}

	}

}
