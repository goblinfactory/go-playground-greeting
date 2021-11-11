package money

// GBP money
type GBP struct {
	Money
}

// NewGBP returns a NewGBP money
func NewGBP(amount float64) GBP {
	m := new2DecimalFormatLeft("£", amount)
	return GBP{m}
}

// Multiply ...(this is an awful hack, just ok for demo's only, do not use in production)
// 3.1 x 1.1 = 31 x 11 / 100 = 341 / 100 = 3.41
func (m GBP) Multiply(amount float64) GBP {
	return NewGBP(m.ToFloat64() * amount)
}

// Add ...
func (m GBP) Add(rhs GBP) GBP {
	return GBP{m.Clone(m.allDigits + rhs.allDigits)}
}

// Subtract ...
func (m GBP) Subtract(rhs GBP) GBP {
	return GBP{m.Money.Clone(m.allDigits - rhs.allDigits)}
}

// GreaterOrEqual ...
func (m GBP) GreaterOrEqual(rhs GBP) bool {
	return m.allDigits >= rhs.allDigits
}

// IsNegative ...
func (m GBP) IsNegative() bool {
	return m.allDigits <= 0
}

// GetMoney ...
func (m GBP) GetMoney() Money {
	return m.Money
}
