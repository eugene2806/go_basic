package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItem(t *testing.T) {
	//1.
	actualCart := Cart{Items: make(map[string]CardItem, 3)}

	productsForAdd := []*Item{
		{"Milk", 1000, 70},
		{"Bread", 200, 40},
		{"Snack", 100, 50},
	}

	for _, product := range productsForAdd {
		actualCart.Items[product.name] = CardItem{
			item:      product,
			weightSum: product.weight,
			priceSum:  product.priceRUB,
			countItem: 1,
		}
	}

	expectedCart := Cart{Items: make(map[string]CardItem, 3)}

	for _, product := range productsForAdd {
		expectedCart.AddItem(product)
	}

	assert.Equal(t, actualCart, expectedCart)

	//2.
	actualCart2 := Cart{Items: make(map[string]CardItem, 3)}

	productsForAddActualCart2 := []*Item{
		{"Milk", 1000, 70},
		{"Bread", 200, 40},
		{"Snack", 100, 50},
	}

	for _, product := range productsForAddActualCart2 {
		actualCart2.Items[product.name] = CardItem{
			item:      product,
			weightSum: product.weight * 2,
			priceSum:  product.priceRUB * 2,
			countItem: 2,
		}
	}

	productsForAddExpectedCart2 := []*Item{
		{"Milk", 1000, 70},
		{"Bread", 200, 40},
		{"Snack", 100, 50},
		{"Milk", 1000, 70},
		{"Bread", 200, 40},
		{"Snack", 100, 50},
	}

	expectedCart2 := Cart{Items: make(map[string]CardItem, 3)}

	for _, product := range productsForAddExpectedCart2 {
		expectedCart2.AddItem(product)
	}

	assert.Equal(t, actualCart2, expectedCart2)
}

func TestRemoveItem(t *testing.T) {

	//Подготовка
	cart := Cart{Items: make(map[string]CardItem, 1)}
	productForDelete := Item{
		name:     "Milk",
		weight:   1000,
		priceRUB: 70,
	}
	cart.Items["Milk"] = CardItem{
		item:      &productForDelete,
		weightSum: 0,
		priceSum:  0,
		countItem: 0,
	}
	//Выполнение
	cart.RemoveItem(&productForDelete)

	//Проверка
	assert.Equal(t, Cart{Items: make(map[string]CardItem)}, cart)

	//Подготовка
	cart2 := Cart{Items: make(map[string]CardItem, 1)}

	productForDelete2 := Item{
		name:     "Milk",
		weight:   1000,
		priceRUB: 70,
	}

	//Выполнение
	cart2.RemoveItem(&productForDelete2)

	//Проверка
	assert.Equal(t, Cart{Items: make(map[string]CardItem)}, cart2)

	//Подготовка
	cart3 := Cart{Items: make(map[string]CardItem, 1)}

	productForDelete3 := Item{
		name:     "Milk",
		weight:   1000,
		priceRUB: 70,
	}
	cart3.Items["Milk"] = CardItem{
		item:      &productForDelete3,
		weightSum: 2000,
		priceSum:  140,
		countItem: 2,
	}

	expectedCart := Cart{Items: make(map[string]CardItem)}

	expectedCart.Items["Milk"] = CardItem{
		item:      &productForDelete3,
		weightSum: 1000,
		priceSum:  70,
		countItem: 1,
	}
	//Выполнение
	cart3.RemoveItem(&productForDelete3)

	//Проверка
	assert.Equal(t, expectedCart, cart3)
}
