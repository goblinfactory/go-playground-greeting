package money

// ZAR South African rand currency
type ZAR struct{ Money }

// NewZAR returns a NewZAR money
func NewZAR(amount float64) ZAR {
	m := new2DecimalFormatLeft("R", amount)
	return ZAR{m}
}

// GetMoney ...
func (m ZAR) GetMoney() Money {
	return m.Money
}
