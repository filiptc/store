package rules

import (
	"github.com/filiptc/store/model"
)

const VOUCHERS_NECESSARY_FOR_DISCOUNT = 2

type voucherRule struct{}

func NewVoucherRule() voucherRule {
	return voucherRule{}
}

func (voucherRule) IsApplicable(cart *model.Cart, item *model.Item) bool {
	if item.Code != "VOUCHER" {
		return false
	}
	vouchersInCart := cart.CountItemsByProductCode("VOUCHER")
	freeVouchersInCart := countFreeVouchers(cart)
	dicountPendingvouchers := vouchersInCart - freeVouchersInCart*VOUCHERS_NECESSARY_FOR_DISCOUNT

	return dicountPendingvouchers >= VOUCHERS_NECESSARY_FOR_DISCOUNT
}

func (voucherRule) Apply(cart *model.Cart) {
	for i, item := range cart.Items {
		if item.Code == "VOUCHER" && item.Price != 0. {
			cart.Items[i].Price = 0
			return
		}
	}
}

func countFreeVouchers(cart *model.Cart) int {
	freeVoucherAmount := 0
	for _, item := range cart.Items {
		if item.Code == "VOUCHER" && item.Price == 0. {
			freeVoucherAmount++
			continue
		}
	}

	return freeVoucherAmount
}
