package money

// GBP money
type GBP struct{ Money }

// NewGBP returns a NewGBP money
func NewGBP(amount float64) GBP {
	m := new2DecimalFormatLeft("£", amount)
	return GBP{m}
}

// Add ...
func (m GBP) Add(rhs GBP) Money {
	return m.Clone(m.allDigits + rhs.allDigits)
}

// Subtract ...
func (m GBP) Subtract(rhs GBP) GBP {
	return GBP{m.Money.Clone(m.allDigits - rhs.allDigits)}
}

// GreaterOrEqual ...
func (m GBP) GreaterOrEqual(rhs GBP) bool {
	return m.allDigits >= rhs.allDigits
}

// GetMoney ...
func (m GBP) GetMoney() Money {
	return m.Money
}
