package model

type Item struct {
	productName     string
	productWeight   int
	productPriceRUB int
}

type Cart struct {
	Items map[string]CardItem
}

type CardItem struct {
	item      *Item
	weightSum int
	priceSum  int
	countItem int
}

func NewItem(productName string, newWeight, newPriceRUB int) *Item {
	return &Item{productName: productName, productWeight: newWeight, productPriceRUB: newPriceRUB}
}

func NewCartItem(item *Item) *CardItem {
	return &CardItem{
		item:      item,
		weightSum: item.productWeight,
		priceSum:  item.productPriceRUB,
		countItem: 1,
	}
}

func NewCart() *Cart {
	return &Cart{Items: make(map[string]CardItem)}
}

func (c *Cart) AddItem(item *Item) {
	nameItem, ok := c.Items[item.productName]

	if !ok {
		c.Items[item.productName] = *NewCartItem(item)

		return
	}

	nameItem.countItem++
	nameItem.priceSum = nameItem.item.productPriceRUB * nameItem.countItem
	nameItem.weightSum = nameItem.item.productWeight * nameItem.countItem
	c.Items[item.productName] = nameItem
}

func (c *Cart) RemoveItem(item *Item) {
	_, ok := c.Items[item.productName]
	if !ok {

		return
	}

	delete(c.Items, item.productName)
}
