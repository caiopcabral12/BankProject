package accounts

import (
	"BankProject/costumers"
)

type BankData struct {
	Costumer      costumers.Costumers
	AccountNumber int
	AgencyNumber  int
	balance       float64
}

func (c *BankData) seeBalance() float64 {
	return c.balance
}

func (c *BankData) Withdraw(withdrawValue float64) (string, float64) {
	possibleWithdraw := c.balance >= withdrawValue && withdrawValue > 0

	if possibleWithdraw {
		c.balance -= withdrawValue
		return "Withdraw DONE! New balance = ", c.balance
	} else {
		return "WITHDRAW REJECTED, Can't withdraw a null value! balance = ", c.balance
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
