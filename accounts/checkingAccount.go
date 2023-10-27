package accounts

import (
	"BankProject/costumers"
	"fmt"
)

type BankData struct {
	Costumer      costumers.Costumers
	AccountNumber int
	AgencyNumber  int
	balance       float64
}

func (c *BankData) Withdraw(withdrawValue float64) string {
	possibleWithdraw := c.balance >= withdrawValue && withdrawValue > 0

	if possibleWithdraw {
		c.balance -= withdrawValue
		return "Withdraw DONE! New balance = "
	} else {
		return "WITHDRAW REJECTED, Can't withdraw a null value! "
	}
}

func (c *BankData) Deposit(depositValue float64) (string, float64) {
	possibleDeposite := depositValue > 0

	if possibleDeposite {
		c.balance += depositValue
		return "Deposite DONE!  New balance =", c.balance
	} else {
		return "Deposite REJECTED, Can't Deposite a null value!  balance =", c.balance
	}
}

func (c *BankData) Transfer(transferValue float64, recipient *BankData) (string, float64) {

	if transferValue <= c.balance && transferValue > 0 {
		c.balance -= transferValue
		recipient.Deposit(transferValue)
		return "Transfer DONE!  New balance =", c.balance
	} else {
		return "Transfer REJECTED, Insuficient balance!  balance =", c.balance
	}

}

func (c *BankData) SeeBalance() float64 {
	return c.balance
}

func (c *BankData) PayBills(accounts verifyAccount, billValue float64) {
	allowPayment := c.balance >= billValue && billValue > 0

	if allowPayment {
		accounts.Withdraw(billValue)
		fmt.Println("Payment done!")
	} else {
		fmt.Println("PAYMENT REJECTED, NOT ENOUGHT BALANCE! ")
	}

}

type verifyAccount interface {
	Withdraw(value float64) string
}
