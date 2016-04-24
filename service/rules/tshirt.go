package rules

import (
	"github.com/filiptc/store/model"
)

const TSHIRTS_NECESSARY_FOR_DISCOUNT = 3

type tshirtRule struct{}

func NewTshirtRule() tshirtRule {
	return tshirtRule{}
}

func (tshirtRule) IsApplicable(cart *model.Cart, item *model.Item) bool {
	return item.Code == "TSHIRT" && hasEnoughTshirtsForDiscount(cart)
}

func (tshirtRule) Apply(cart *model.Cart) {
	for i, item := range cart.Items {
		if item.Code == "TSHIRT" {
			cart.Items[i].Price = 19.
		}
	}
}

func hasEnoughTshirtsForDiscount(cart *model.Cart) bool {
	return cart.CountItemsByProductCode("TSHIRT") >= TSHIRTS_NECESSARY_FOR_DISCOUNT
}
