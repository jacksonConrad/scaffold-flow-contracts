package test_main

import (
	"testing"

	"github.com/bjartek/overflow/overflow"
)

type OverflowTestUtils struct {
	T *testing.T
	O *overflow.Overflow
}

const ServiceAddress = "0xf8d6e0586b0a20c7"

func NewOverflowTest(t *testing.T) *OverflowTestUtils {
	return &OverflowTestUtils{T: t, O: overflow.NewTestingEmulator().Start()}
	// return &OverflowTestUtils{T: t, O: overflow.NewOverflowEmulator().Start()}
}

func (otu *OverflowTestUtils) HelloWorldScript() string {
	val, _ := otu.O.ScriptFromFile("test").
		Args(otu.O.Arguments()).
		RunReturns()

	return val.ToGoValue().(string)
}

func (otu *OverflowTestUtils) HelloWorldTx() *OverflowTestUtils {
	otu.O.TransactionFromFile("test").
		SignProposeAndPayAs("account").
		Args(otu.O.Arguments()).
		Test(otu.T).
		AssertSuccess()

	return otu
}
