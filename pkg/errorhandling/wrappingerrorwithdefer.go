package errorhandling

import (
	"errors"
	"fmt"

	"github.com/goblinfactory/greeting/pkg/money"
)

// DemoWrappingErrorsUsingDefer shows how to add the same error wrapping behavior to all returns using defer
func DemoWrappingErrorsUsingDefer(num int) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Number error: %w", err)
		}
	}()

	if num < 0 {
		return 0, errors.New("no negative numbers allowed")
	}

	if num%2 == 0 {
		return 0, errors.New("no even numbers allowed")
	}
	if num%2 == 1 {
		return 0, errors.New("No odd numbers allowed")
	}

	return 10, nil
}

type account struct {
	Locked  bool
	Balance money.GBP
}

func closeAccount(a account) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Account error: %w", err)
		}
	}()

	if a.Balance.IsNegative() {
		return errors.New("cannot close account with a negative balance")
	}
	if a.Locked {
		return errors.New("cannot close a locked account. Please unlock the account first")
	}

	return nil
}
