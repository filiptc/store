package store

import (
	"testing"

	"github.com/filiptc/store/service"
	"github.com/filiptc/store/service/rules"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type FunctionalSuite struct{}

var _ = Suite(&FunctionalSuite{})

type expectedResults struct {
	productCodes []string
	price        float64
}

var expected = []expectedResults{
	{
		[]string{"VOUCHER", "TSHIRT", "MUG"},
		32.5,
	}, {
		[]string{"VOUCHER", "TSHIRT", "VOUCHER"},
		25,
	}, {
		[]string{"TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"},
		81,
	}, {
		[]string{"VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"},
		74.5,
	},
}

func (s *FunctionalSuite) TestImplementation(c *C) {
	pricingRules := []rules.Rule{
		rules.NewVoucherRule(),
		rules.NewTshirtRule(),
	}
	for _, e := range expected {
		co := service.NewCheckout(pricingRules)
		for _, productCode := range e.productCodes {
			co.Scan(productCode)
		}
		c.Assert(co.GetTotal(), Equals, e.price)
	}
}
