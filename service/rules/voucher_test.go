package rules

import (
	"github.com/filiptc/store/model"
	. "gopkg.in/check.v1"
)

type VoucherRuleSuite struct {
	rule voucherRule
}

var _ = Suite(&VoucherRuleSuite{})

func (s *VoucherRuleSuite) SetUpSuite(c *C) {
	s.rule = NewVoucherRule()
}

func (s *VoucherRuleSuite) TestIsApplicable(c *C) {
	cart := &model.Cart{}

	voucher, _ := model.NewItem("VOUCHER")
	c.Assert(s.rule.IsApplicable(cart, voucher), Equals, false)

	cart.AddItem(newVoucher())
	cart.AddItem(newVoucher())
	c.Assert(s.rule.IsApplicable(cart, voucher), Equals, true)

	cart.Items[0].Price = 0.
	c.Assert(s.rule.IsApplicable(cart, voucher), Equals, false)

	cart.AddItem(newVoucher())
	c.Assert(s.rule.IsApplicable(cart, voucher), Equals, false)

	cart.AddItem(newVoucher())
	c.Assert(s.rule.IsApplicable(cart, voucher), Equals, true)

	tshirt, _ := model.NewItem("TSHIRT")
	c.Assert(s.rule.IsApplicable(cart, tshirt), Equals, false)
}

func (s *VoucherRuleSuite) TestCountFreeVouchers(c *C) {
	item1, _ := model.NewItem("VOUCHER")
	item2, _ := model.NewItem("VOUCHER")
	item3, _ := model.NewItem("VOUCHER")
	cart := model.NewCart(item1, item2, item3)
	c.Assert(countFreeVouchers(cart), Equals, 0)
	cart.Items[0].Price = 0.
	c.Assert(countFreeVouchers(cart), Equals, 1)
}

func (s *VoucherRuleSuite) TestApply(c *C) {
	cart := model.NewCart(newVoucher(), newVoucher())
	s.rule.Apply(cart)
	c.Assert(cart.Items[0].Price, Equals, 0.)
}

func newVoucher() *model.Item {
	voucher, _ := model.NewItem("VOUCHER")
	return voucher
}
