package testaccountservice

import (
	"fmt"

	"github.com/goblinfactory/greeting/pkg/accountservice"
	"github.com/goblinfactory/greeting/pkg/money"
)

// TestAccountService ...
func TestAccountService() {
	as := accountservice.New("accounts.json")
	as.CreateAccount("H1", "Harrys account", money.NewGBP(0.01))
	as.CreateAccount("H2", "Haggard account", money.NewGBP(4500))
	as.CreateAccount("N1", "Nelly account1", money.NewZAR(18325.01))
	as.CreateAccount("N2", "Nelly account2", money.NewZAR(18325.99))
	as.CreateAccount("C1", "Chandler account", money.NewGBP(50.95))
	as.Dump()
	fmt.Println("---")
	as.Transfer("H2", "H1", money.NewGBP(0.01))
	as.Dump()
}
