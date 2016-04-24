package service

import (
	"github.com/filiptc/store/model"
	"github.com/filiptc/store/service/rules"
	. "gopkg.in/check.v1"
)

type CheckoutSuite struct{}

var _ = Suite(&CheckoutSuite{})

func getService(c *C) *Checkout {
	service := NewCheckout([]rules.Rule{})
	c.Assert(service, NotNil)
	return service
}

func (s *CheckoutSuite) TestScan(c *C) {
	service := getService(c)
	service.Scan("VOUCHER")
	service.wg.Wait()
	c.Assert(len(service.cart.Items), Equals, 1)

}

func (s *CheckoutSuite) TestGetTotal(c *C) {
	service := getService(c)
	service.cart.Items = []*model.Item{
		{"foo", "bar", 10},
		{"foo", "bar", 10},
	}
	c.Assert(service.GetTotal(), Equals, 20.)
}

func (s *CheckoutSuite) TestConsumeScans(c *C) {
	service := getService(c)
	service.wg.Add(1)
	go func() { service.consumeScans() }()
	service.scans <- &model.Item{"foo", "bar", 10}
	service.wg.Wait()
	c.Assert(len(service.cart.Items), Equals, 1)
}
