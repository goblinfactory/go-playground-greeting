package accountservice

import (
	"fmt"
	"sort"
	"sync"
)

// Accounts ...
type Accounts map[ID]Account

type inMemoryDb struct {
	l sync.RWMutex
	Accounts
}

func newInMemoryDb(filename string) *inMemoryDb {
	return &inMemoryDb{Accounts: Accounts{}}
}

func (db *inMemoryDb) Create(account Account) {
	db.l.Lock()
	defer db.l.Unlock()
	db.Accounts[account.ID] = account
}

func (db *inMemoryDb) Read(id ID) Account {
	db.l.RLock()
	defer db.l.RUnlock()
	return db.Accounts[id]
}

func (db *inMemoryDb) Dump() {
	db.l.RLock()
	defer db.l.RUnlock()

	// do a small dance to sort the keys
	keys := make([]string, 0, len(db.Accounts))
	for k := range db.Accounts {
		keys = append(keys, string(k))
	}

	sort.Strings(keys)

	// in C# we'd do for (v in keys.OrderBy(...))
	for _, k := range keys {
		v := db.Accounts[ID(k)]
		fmt.Println(k, v.Balance.Wide(), ":", v.Balance.CurrencyDescription())
	}
}

// Update locks the database for reading and uses optimistic concurrency
// to check that the existing account balances are all correct
// then locks the database for writing and pretends to updates all the accounts
// in a single atomic transaction. Sort of, good enough for demo purposes.
func (db *inMemoryDb) Update(updates ...Update) (bool, error) {
	db.l.Lock()
	defer db.l.Unlock()
	for _, u := range updates {
		b := db.Accounts[u.ID].Balance
		if b != u.OldBalance {
			return false, fmt.Errorf("expected account[%s] balance to be %s, but found %s. Operation cancelled. (Optimistic lock error)", u.ID, b.Wide(), u.OldBalance.Wide())
		}
	}
	for _, u := range updates {
		account := db.Accounts[u.ID]
		account.Balance = u.NewBalance
		db.Accounts[u.ID] = account
	}
	return true, nil
}
