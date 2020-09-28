package decimalflag

import "github.com/shopspring/decimal"

// DecimalFlag implements standard library flag for
// github.com/shopspring/decimal.
type DecimalFlag decimal.Decimal

// String implements flag.Value interface
func (d *DecimalFlag) String() string {
	return decimal.Decimal(*d).String()
}

// Set implements flag.Value interface
func (d *DecimalFlag) Set(value string) error {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return err
	}
	*d = DecimalFlag(v)
	return nil
}

// Get implments flag.Getter interface.
func (d *DecimalFlag) Get() interface{} {
	return decimal.Decimal(*d)
}

// Decimal returns the decimal.Decimal value
func (d *DecimalFlag) Decimal() decimal.Decimal {
	return decimal.Decimal(*d)
}
