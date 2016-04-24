package rules

import "github.com/filiptc/store/model"

type Rule interface {
	IsApplicable(*model.Cart, *model.Item) bool
	Apply(*model.Cart)
}
