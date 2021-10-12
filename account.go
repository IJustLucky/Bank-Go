package main

import (
	"errors"
	"time"
)

type account struct {
	name string
	txs  []interface{}
}

type deposit struct {
	amount float64
	time   time.Time
}

type withdraw struct {
	amount float64
	time   time.Time
}

//a point to account is made because it holds txs
//returns moneyz- a float 64
//when deposit or withdraw is called- a type- the the amount is either added or taken away
func (a *account) balance() (moneyz float64) {
	//this will return the totality of all transactions
	for _, tx := range a.txs {
		//checks if the tx is type deposit or withdrawal
		switch tx := tx.(type) {
		case deposit:
			moneyz += tx.amount
		case withdraw:
			moneyz -= tx.amount
		}
	}
	return
}

//returns an error if the amount in is less than or equal to 0
func (a *account) deposit(in float64) error {
	if in <= 0 {
		return errors.New("can only deposit positive values")
	}
	//d is now type deposit and has parameters in and time.Now
	d := deposit{
		amount: in,
		time:   time.Now(),
	}
	//this appends the amount and time that the transaction goes on to a.txs and returns nothing
	a.txs = append(a.txs, d)
	return nil
}

//error if withdraw is less than or equal to 0
func (a *account) withdraw(out float64) error {
	if out <= 0 {
		return errors.New("can only withdraw positive values")
	}
	//error if you try to withdraw more than is in the balance
	if out > a.balance() {
		return errors.New("insufficient funds")
	}
	//same as the deposit type except this is withdrawal not balance
	w := withdraw{
		amount: out,
		time:   time.Now(),
	}
	a.txs = append(a.txs, w)
	return nil
}

func (a *account) MarshalJSON() ([]byte, error) {
	return nil, nil
}
