package accounts

import (
	"BankProject/costumers"
	"fmt"
)

type Savings struct {
	Costumer      costumers.Costumers
	AccountNumber int
	AgencyNumber  int
	Operation     int
	balance       float64
}

func (c *Savings) Withdraw(withdrawValue float64) (string, float64) {
	possibleWithdraw := c.balance >= withdrawValue && withdrawValue > 0

	if possibleWithdraw {
		c.balance -= withdrawValue
		return "Withdraw DONE! New balance On savings = ", c.balance
	} else {
		return "WITHDRAW REJECTED, Can't withdraw a null value! ", withdrawValue
	}
}

func (c *Savings) Deposit(depositValue float64) (string, float64) {
	possibleDeposite := depositValue > 0

	if possibleDeposite {
		c.balance += depositValue
		return "Deposite DONE!  Value = ", depositValue
	} else {
		return "Deposite REJECTED, Can't Deposite a null value!", depositValue
	}
}

func (c *Savings) SeeBalance() float64 {
	return c.balance
}

func (c *Savings) PayBills(accounts verifyAccount, billValue float64) {
	allowPayment := c.balance >= billValue && billValue > 0

	if allowPayment {
		accounts.Withdraw(billValue)
		fmt.Println("Payment done!")
	} else {
		fmt.Println("PAYMENT REJECTED, NOT ENOUGHT BALANCE! ")
	}

}
