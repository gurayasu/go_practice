package bank

import "testing"

func TestAccount(t *testing.T){
	account := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error("cannot create an Account object")
	}
}

func TestDeposit(t *testing.T){
	account := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}
	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T){
	account := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}
	
	if err := account.Deposit(-10); err == nil {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestWithdraw(t *testing.T){
	account := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T){
	account := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()

	if statement != "1001 - John - 100"{
		t.Error("statementがフォーマットに沿ってないよ")
	}
}

func TestTransfer(t *testing.T){
	accountA := Account{
		Customer: Customer{
			Name: "John",
			Address: "Japan, Kyoto",
			Phone: "060888999",
		},
		Number: 1001,
		Balance: 0,
	}

	accountB := Account{
		Customer: Customer{
			Name:    "Mark",
			Address: "Irvine, California",
			Phone:   "(949) 555 0198",
		},
		Number:  1002,
		Balance: 0,
	}

	accountA.Deposit(100)
	err := accountA.Transfer(50, &accountB)

	if accountA.Balance != 50 && accountB.Balance != 50 {
		t.Error("accountA -> accountA transfer is not working", err)
	}
}