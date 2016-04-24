package model

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CartSuite struct {
	cart *Cart
}

var _ = Suite(&CartSuite{})

func (s *CartSuite) SetUpTest(c *C) {
	s.cart = NewCart()
}

func (s *CartSuite) TestGetItemPriceSum(c *C) {
	voucher, err := NewItem("VOUCHER")
	c.Assert(err, IsNil)
	s.cart.AddItem(voucher)

	tshirt, err := NewItem("TSHIRT")
	c.Assert(err, IsNil)
	s.cart.AddItem(tshirt)

	mug, err := NewItem("MUG")
	c.Assert(err, IsNil)
	s.cart.AddItem(mug)

}

func (s *CartSuite) TestCountItemsByProductCode(c *C) {
	voucher, err := NewItem("VOUCHER")
	c.Assert(err, IsNil)
	mug, err := NewItem("MUG")
	c.Assert(err, IsNil)

	s.cart.AddItem(voucher)
	s.cart.AddItem(voucher)
	s.cart.AddItem(voucher)
	s.cart.AddItem(mug)

	c.Assert(s.cart.CountItemsByProductCode("VOUCHER"), Equals, 3)
	c.Assert(s.cart.CountItemsByProductCode("TSHIRT"), Equals, 0)
	c.Assert(s.cart.CountItemsByProductCode("MUG"), Equals, 1)
}
