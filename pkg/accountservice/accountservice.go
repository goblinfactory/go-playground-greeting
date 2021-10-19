package accountservice

import (
	"log"

	"github.com/goblinfactory/greeting/pkg/money"
)

// Update ...
type Update struct {
	ID
	OldBalance money.Money // Old balance so that you can use optimistic locking to ensure it hasnt changed.
	NewBalance money.Money // New balance to set account to.
}

// ReaderWriter ...
type ReaderWriter interface {
	Create(account Account)
	Read(id ID) Account
	Update(updates ...Update) (bool, error)
	Dump()
}

// Account ...
type Account struct {
	ID
	Name    string
	Balance money.Money
}

// ID of the account
type ID string

// AnyCurrency ...
type AnyCurrency interface {
	GetMoney() money.Money
}

// AccountService manages accounts
type AccountService struct {
	db ReaderWriter
}

// New returns a new AccountService
func New(filename string) AccountService {
	db := newInMemoryDb(filename)
	return AccountService{db}
}

// CreateAccount creates a new account (does not check if already exists) and returns account added to db
func (s *AccountService) CreateAccount(id string, name string, startingBalance AnyCurrency) Account {
	accID := ID(id)
	ac := Account{accID, name, startingBalance.GetMoney()}
	s.db.Create(ac)
	return ac
}

// Transfer funds
func (s *AccountService) Transfer(fromAccountID string, toAccountID string, amount AnyCurrency) (money.Money, money.Money) {
	from := s.db.Read(ID(fromAccountID))
	to := s.db.Read(ID(toAccountID))

	if from.Balance.CurrencySymbol != to.Balance.CurrencySymbol {
		log.Fatalf("Mismatched currencies. %s != %s", from.Balance.CurrencySymbol, to.Balance.CurrencySymbol)
	}

	if !from.Balance.GreaterOrEqual(to.Balance) {
		log.Fatalf("Insufficient funds. From:[%s]:%s, To:[%s]:%s", from.ID, from.Balance.Wide(), to.ID, to.Balance.Wide())
	}

	newFromBalance := from.Balance.Subtract(amount.GetMoney())
	newToBalance := to.Balance.Add(amount.GetMoney())

	update1 := Update{from.ID, from.Balance, newFromBalance}
	update2 := Update{to.ID, to.Balance, newToBalance}

	if success, error := s.db.Update(update1, update2); !success {
		// no retries in this simple version, can add that later.
		// depending on what is making calls to the service, retries may be necessary.
		log.Fatal(error)
	}
	return newFromBalance, newToBalance
}

// Dump prints out the state of the service and dependencies
func (s *AccountService) Dump() {
	s.db.Dump()
}
