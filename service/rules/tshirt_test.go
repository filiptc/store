package rules

import (
	"github.com/filiptc/store/model"
	. "gopkg.in/check.v1"
)

type TshirtRuleSuite struct {
	rule tshirtRule
}

var _ = Suite(&TshirtRuleSuite{})

func (s *TshirtRuleSuite) SetUpSuite(c *C) {
	s.rule = NewTshirtRule()
}

func (s *TshirtRuleSuite) TestIsApplicable(c *C) {
	item, _ := model.NewItem("VOUCHER")
	cart := &model.Cart{}
	c.Assert(s.rule.IsApplicable(cart, item), Equals, false)

	voucher, _ := model.NewItem("VOUCHER")
	mug, _ := model.NewItem("MUG")
	tshirt, _ := model.NewItem("TSHIRT")
	cart = &model.Cart{}
	cart.AddItem(voucher)
	cart.AddItem(tshirt)
	cart.AddItem(voucher)
	cart.AddItem(voucher)
	cart.AddItem(mug)
	cart.AddItem(tshirt)
	cart.AddItem(tshirt)
	c.Assert(s.rule.IsApplicable(cart, tshirt), Equals, true)
}

func (s *TshirtRuleSuite) TestHasEnoughTshirtsForDiscount(c *C) {
	tshirt, _ := model.NewItem("TSHIRT")
	cart := &model.Cart{}
	cart.AddItem(tshirt)
	cart.AddItem(tshirt)
	c.Assert(hasEnoughTshirtsForDiscount(cart), Equals, false)

	voucher, _ := model.NewItem("VOUCHER")
	mug, _ := model.NewItem("MUG")
	cart = &model.Cart{}
	cart.AddItem(voucher)
	cart.AddItem(tshirt)
	cart.AddItem(voucher)
	cart.AddItem(voucher)
	cart.AddItem(mug)
	cart.AddItem(tshirt)
	cart.AddItem(tshirt)
	c.Assert(hasEnoughTshirtsForDiscount(cart), Equals, true)
}

func (s *TshirtRuleSuite) TestApply(c *C) {
	tshirt, _ := model.NewItem("TSHIRT")
	cart := model.NewCart(tshirt, tshirt, tshirt)
	s.rule.Apply(cart)
	c.Assert(cart.Items[0].Price, Equals, 19.)
}
