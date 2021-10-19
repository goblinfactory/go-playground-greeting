package money

// USD Dollar currency
type USD struct{ Money }

// Subtract ...
func (m USD) Subtract(rhs USD) USD {
	return USD{m.Money.Clone(m.allDigits - rhs.allDigits)}
}

// NewUSD returns a New US Dollar money
func NewUSD(amount float64) USD {
	m := new2DecimalFormatLeft("$", amount)
	return USD{m}
}

// GetMoney ...
func (m USD) GetMoney() Money {
	return m.Money
}
