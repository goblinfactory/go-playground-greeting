package money

import (
	"fmt"
	"log"
	"strconv"
)

//TODO: for simple hacking and demos, convert this to use float64 as a base type, with only 2 decimal precision.

// Money ...
type Money struct {
	allDigits        int64
	integralDigits   int64
	fractionalDigits int64
	numDecimals      int64
	mod              int
	CurrencySymbol   string
	formatShort      string
	formatWide       string
	formatString     string
	parseString      string
}

// CurrencyNames ...
var CurrencyNames = map[string]string{
	"$": "US Dollar",
	"R": "South African Rands",
	"£": "GBP Sterling",
}

func new2DecimalFormatLeft(currencySymbol string, amount float64) Money {
	allDigits := int64((amount + 0.005) * float64(100))
	fractionalDigits := allDigits % 100
	integralDigits := (allDigits - fractionalDigits) / 100

	m := Money{
		allDigits,
		integralDigits,
		fractionalDigits,
		2,
		100,
		currencySymbol,
		currencySymbol + " %4d.%02d",
		currencySymbol + " %11d.%02d",
		currencySymbol + " %d.%02d",
		"%d.%02d",
	}
	return m
}

// GreaterOrEqual ...
func (m Money) GreaterOrEqual(rhs Money) bool {
	return m.allDigits >= rhs.allDigits
}

// Add ...
func (m Money) Add(rhs Money) Money {
	return m.Clone(m.allDigits + rhs.allDigits)
}

// Multiply ...(this is an awful hack, just ok for demo's only, do not use in production)
func (m Money) Multiply(amount float64) Money {
	fv := m.ToFloat64()
	newValue := fv * amount
	return m.Clone(int64(newValue))
}

// Subtract ...
func (m Money) Subtract(rhs Money) Money {
	return m.Clone(m.allDigits - rhs.allDigits)
}

// Clone ...
func (m Money) Clone(allDigits int64) Money {

	fractionalDigits := allDigits % 100
	integralDigits := (allDigits - fractionalDigits) / 100
	return Money{
		allDigits,
		integralDigits,
		fractionalDigits,
		2,
		100,
		m.CurrencySymbol,
		m.formatShort,
		m.formatWide,
		m.formatString,
		m.parseString,
	}
}

// Short returns a short format, typically 7 characters wide 0000.00
func (m Money) Short() string {
	return fmt.Sprintf(m.formatShort, m.integralDigits, m.fractionalDigits)
}

// ToFloat64 returns a float64
func (m Money) ToFloat64() float64 {
	text := fmt.Sprintf(m.parseString, m.integralDigits, m.fractionalDigits)
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// ToString returns a formatted string with currency symbol and currency amount with fractional digits
func (m Money) ToString() string {
	return fmt.Sprintf(m.formatString, m.integralDigits, m.fractionalDigits)
}

// Wide returns a wide format,  14  characters wide,
func (m Money) Wide() string {
	return fmt.Sprintf(m.formatWide, m.integralDigits, m.fractionalDigits)
}

// CurrencyDescription ...
func (m Money) CurrencyDescription() string {
	return CurrencyNames[m.CurrencySymbol]
}

// SampleUsage ...
func SampleUsage() {
	usd1 := NewUSD(100.123)
	gbp1 := NewGBP(500)
	zar1 := NewZAR(123.45)
	zar2 := NewZAR(123.99)
	fmt.Println(usd1.Short(), usd1.CurrencyDescription())
	fmt.Println(gbp1.Short(), gbp1.CurrencyDescription())
	fmt.Println(zar1.Short(), zar1.CurrencyDescription())
	fmt.Println(zar2.Short(), zar2.CurrencyDescription())
}

// references
// https://floating-point-gui.de/formats/integer/   : see the note and warning NOT to use integers if at all possible.
