package model

import "fmt"

type ProductCode string

const ITEM_NOT_FOUND = "Item with code %v not found"

type Item struct {
	Code  string
	Name  string
	Price float64
}

var Items = map[string]Item{
	"VOUCHER": Item{"VOUCHER", "Cabify Voucher", 5},
	"TSHIRT":  Item{"TSHIRT", "Cabify T-Shirt", 20},
	"MUG":     Item{"MUG", "Cabify Coffee Mug", 7.5},
}

func NewItem(code string) (*Item, error) {
	if i, ok := Items[code]; ok {
		return &i, nil
	}
	return nil, fmt.Errorf(ITEM_NOT_FOUND, code)
}
