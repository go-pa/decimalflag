package decimalflag

import (
	"bytes"
	"flag"
	"testing"

	"github.com/shopspring/decimal"
)

func TestParseOK(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	var df DecimalFlag
	fs.Var(&df, "v", "")
	const num = "-1.2345678911111111112222222222333333333344444444445555555555"
	err := fs.Parse([]string{"-v", num})
	if err != nil {
		t.Fatal(err)
	}
	if df.Decimal().String() != num {
		t.Fatal(df.Decimal().String(), num)

	}

	vv, ok := df.Get().(decimal.Decimal)
	if !ok {
		t.Fatal(vv)
	}
	if !df.Decimal().Equal(vv) {
		t.Fatal(df.Decimal().String(), num)
	}
}

func TestParseError(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	b := bytes.NewBuffer([]byte{})
	fs.SetOutput(b)
	var df DecimalFlag
	fs.Var(&df, "v", "")

	err := fs.Parse([]string{"-v", "aslkdjs"})
	if err == nil {
		t.Fatal("expected error")
	}

	if df.Decimal().String() != "0" {
		t.Fatal(df.Decimal().String())

	}
}
