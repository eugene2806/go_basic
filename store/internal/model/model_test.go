package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItem(t *testing.T) {
	cart := NewCart()
	productsCase1 := []*Item{
		NewItem("Milk", 1000, 80),
		NewItem("Bread", 200, 40),
		NewItem("Snack", 80, 50),
	}

	for _, item := range productsCase1 {
		_, ok := cart.Items[item.productName]

		assert.Equal(t, false, ok, "Товар уже есть в корзине")
		cart.AddItem(item)

		_, ok = cart.Items[item.productName]
		assert.Equal(t, true, ok, "Товара нет в корзине")
	}

	productsCase2 := []*CardItem{
		&CardItem{item: NewItem("Milk", 1000, 80),
			weightSum: 2000,
			priceSum:  160,
			countItem: 2},
		&CardItem{item: NewItem("Bread", 200, 40),
			weightSum: 400,
			priceSum:  80,
			countItem: 2},
		&CardItem{item: NewItem("Snack", 80, 50),
			weightSum: 160,
			priceSum:  100,
			countItem: 2},
	}

	for _, item := range productsCase1 {
		nameItem, _ := cart.Items[item.productName]

		nameItem.countItem++
		nameItem.priceSum = nameItem.item.productPriceRUB * nameItem.countItem
		nameItem.weightSum = nameItem.item.productWeight * nameItem.countItem

		cart.Items[item.productName] = nameItem

	}

	for _, cartItem := range productsCase2 {
		actualCount := cart.Items[cartItem.item.productName].countItem
		actualPriceSum := cart.Items[cartItem.item.productName].priceSum
		actualWeightSum := cart.Items[cartItem.item.productName].weightSum

		assert.Equal(t, cartItem.countItem, actualCount, "Количество товаров не совпадает")
		assert.Equal(t, cartItem.priceSum, actualPriceSum, "Суммы за товар не совпадают")
		assert.Equal(t, cartItem.weightSum, actualWeightSum, "Вес товаров не совпадает")
	}

}

func TestRemoveItem(t *testing.T) {
	cart := NewCart()
	product1 := NewItem("Milk", 1000, 80)
	cart.AddItem(product1)

	_, isOk := cart.Items[product1.productName]
	assert.Equal(t, true, isOk, "Ожидалось наличие товара в корзине")
	cart.RemoveItem(product1)

	_, isOk = cart.Items[product1.productName]
	assert.Equal(t, false, isOk, "Ожидалось отсутствие товара в корзине")
}
