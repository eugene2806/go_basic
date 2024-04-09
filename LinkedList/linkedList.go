package main

import "fmt"

var uniqueValues = make(map[int]int, 40)

type Element struct {
	data int
	next *Element
}

func genNewList(data int) *Element {
	uniqueValues[data] = data

	return &Element{data: data, next: nil}
}

func detectDuplicate(data int) bool {
	if _, ok := uniqueValues[data]; ok {

		return true
	}

	return false
}

func (el *Element) addToEnd(data int) {
	currEl := el
	var lastEl *Element

	for {
		if currEl.next != nil {

			currEl = currEl.next
			continue
		}

		if detectDuplicate(data) {
			lastEl = currEl

			return
		}

		lastEl = genNewList(data)
		currEl.next = lastEl

		return
	}

}

func (el *Element) addToHead(data int) *Element {
	currEl := el
	var newEl *Element

	if detectDuplicate(data) {

		return currEl
	}

	newEl = genNewList(data)

	newEl.next = currEl
	currEl = newEl

	return newEl
}

func (el *Element) addToIndex(index int, dataList ...int) {
	countIndex := 0
	currEl := el
	var newEl *Element

	for _, data := range dataList {

		if detectDuplicate(data) {
			continue
		}

		newEl = genNewList(data)

		for {
			if countIndex != index && currEl.next != nil {
				currEl = currEl.next
				countIndex++
				continue
			}

			if currEl.next == nil {
				fmt.Println("Для добавления элемента в конец списка используйте - addToEnd()")

				return
			}

			newEl.next = currEl.next
			currEl.next = newEl
			index++
			break
		}
	}
}

func (el *Element) removeToEnd() {
	currEl := el
	var prevEl *Element

	for {
		if currEl.next != nil {
			prevEl = currEl
			currEl = currEl.next
			continue
		}

		delete(uniqueValues, currEl.data)
		prevEl.next = nil

		return
	}
}

func (el *Element) removeToHead() *Element {
	currEl := el
	var prevEl *Element

	if currEl.next != nil {
		prevEl = currEl
		currEl = currEl.next
		delete(uniqueValues, prevEl.data)
		prevEl.next = nil

		return currEl
	}

	return currEl
}

func (el *Element) removeToIndex(index int) {
	currEl := el
	var prevEl *Element
	countIndex := 0

	if index == 0 {
		return
	}

	for {
		if countIndex != index && currEl.next != nil {
			prevEl = currEl
			currEl = currEl.next
			countIndex++
			continue
		}

		if currEl.next == nil {
			fmt.Println("Для удаления элемента в конеце списка используйте - removeToEnd()")

			return
		}

		delete(uniqueValues, currEl.data)
		prevEl.next = currEl.next

		break
	}
}

func (el *Element) search(data int) (index, dataEl int) {
	currEl := el
	countIndex := 0

	for {
		if currEl.data == data {
			return countIndex, currEl.data
		}

		if currEl.next != nil {
			currEl = currEl.next
			countIndex++
			continue
		}
		break
	}
	fmt.Printf("Элемент => %d не найден\n", data)

	return 0, 0
}

func printList(el *Element) {
	fmt.Printf("[")
	for {
		if el.next == nil {
			fmt.Printf("%d", el.data)
			fmt.Printf("]\n")

			return
		}
		fmt.Printf("%d, ", el.data)

		if el.next != nil {
			el = el.next
			continue
		}
	}
}

func main() {

	head := genNewList(10)
	head.addToEnd(15)

	head = head.addToHead(10)
	head.addToEnd(18)
	fmt.Println("First Element", head)
	head.addToEnd(19)

	head = head.addToHead(0)
	fmt.Println("First Element", head)
	printList(head)

	fmt.Println("===============")

	head.addToIndex(3, 1, 2, 3)
	printList(head)

	head.removeToEnd()
	head.removeToEnd()
	head = head.removeToHead()
	printList(head)

	head.removeToIndex(3)
	printList(head)

	index, val := head.search(15)
	fmt.Printf("Индекс => %d Элемент => %d\n", index, val)

}
