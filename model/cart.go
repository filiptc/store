package model

// struct instead of map[int]Item for future-proofness
type Cart struct {
	Items []*Item
}

func NewCart(items ...*Item) *Cart {
	return &Cart{items}
}

func (c *Cart) AddItem(i *Item) {
	c.Items = append(c.Items, i)
}

func (c *Cart) GetItemPriceSum() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Price
	}
	return total
}

func (c *Cart) CountItemsByProductCode(code string) int {
	var count int
	for _, item := range c.Items {
		if item.Code == code {
			count++
		}
	}
	return count

}
