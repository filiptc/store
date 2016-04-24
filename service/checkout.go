package service

import (
	"sync"

	"github.com/filiptc/store/model"
	"github.com/filiptc/store/service/rules"
)

type Checkout struct {
	rules []rules.Rule
	cart  *model.Cart
	scans chan *model.Item
	wg    sync.WaitGroup
}

func NewCheckout(rules []rules.Rule) *Checkout {
	co := &Checkout{
		rules: rules,
		cart:  model.NewCart(),
		scans: make(chan *model.Item, 1),
	}
	go co.consumeScans()
	return co
}

func (s *Checkout) Scan(code string) error {
	s.wg.Add(1)
	item, err := model.NewItem(code)
	if err != nil {
		return err
	}

	go func() { s.scans <- item }()

	return nil
}

func (s *Checkout) GetTotal() float64 {
	s.wg.Wait()
	return s.cart.GetItemPriceSum()
}

func (s *Checkout) consumeScans() {
	for {
		item := <-s.scans
		s.cart.AddItem(item)
		s.applyRules(item)
		s.wg.Done()
	}
}

func (s *Checkout) applyRules(item *model.Item) {
	for _, rule := range s.rules {
		if rule.IsApplicable(s.cart, item) {
			rule.Apply(s.cart)
		}
	}
}
