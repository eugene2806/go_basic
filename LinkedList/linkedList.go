package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func genNewList(data int) *ListNode {
	return &ListNode{Val: data, Next: nil}
}

func (el *ListNode) Append(data int) {
	currEl := el
	var lastEl *ListNode

	for currEl.Next != nil {
		currEl = currEl.Next
	}

	lastEl = genNewList(data)
	currEl.Next = lastEl
}

func (el *ListNode) Prepend(data int) *ListNode {
	currEl := el
	var newEl *ListNode

	newEl = genNewList(data)

	newEl.Next = currEl
	currEl = newEl

	return newEl
}

func (el *ListNode) Insert(index int, dataList ...int) {
	countIndex := 0
	currEl := el
	var newEl *ListNode

	for _, data := range dataList {

		newEl = genNewList(data)

		for countIndex != index && currEl.Next != nil {
			currEl = currEl.Next
			countIndex++
		}

		if currEl.Next == nil {
			fmt.Println("Для добавления элемента в конец списка используйте - Append()")

			return
		}

		newEl.Next = currEl.Next
		currEl.Next = newEl
		index++
	}
}

func (el *ListNode) Delete(index int) *ListNode {
	firstEl := el
	currEl := el
	var prevEl *ListNode
	countIndex := 0

	if index == 0 {
		currEl = currEl.Next
		firstEl.Next = nil

		return currEl
	}

	for countIndex != index && currEl.Next != nil {
		prevEl = currEl
		currEl = currEl.Next
		countIndex++
	}

	if index > countIndex {
		fmt.Println("Unknown index")
		return firstEl
	}

	prevEl.Next = currEl.Next
	return firstEl
}

func (el *ListNode) Length() int {
	currEl := el
	count := 0

	for currEl != nil {
		currEl = currEl.Next
		count++
	}

	return count
}

func (el *ListNode) Search(data int) (index int, node *ListNode) {
	currEl := el
	firstEl := el
	countIndex := 0

	for currEl != nil {
		if currEl.Val == data {
			return countIndex, currEl
		}

		currEl = currEl.Next
		countIndex++
	}

	fmt.Printf("Элемент => %d не найден\n", data)

	return 0, firstEl
}

func (el *ListNode) Get(index int) *ListNode {
	currEl := el
	firstEl := el
	countIndex := 0

	for index != countIndex && currEl.Next != nil {
		currEl = currEl.Next
		countIndex++
	}

	if index > countIndex {
		fmt.Println("Unknown index")

		return firstEl
	}

	return currEl
}

func (el *ListNode) Update(index, data int) {
	currEl := el
	countIndex := 0

	for index != countIndex && currEl.Next != nil {
		currEl = currEl.Next
		countIndex++
	}

	if index > countIndex {
		fmt.Println("Unknown index")

		return
	}

	currEl.Val = data
}

func (el *ListNode) Traverse() {
	currEl := el

	fmt.Printf("[")

	for currEl.Next != nil {
		fmt.Printf("%d, ", currEl.Val)
		currEl = currEl.Next
	}

	fmt.Printf("%d", currEl.Val)
	fmt.Printf("]\n")
}

func main() {

	head := genNewList(10)

	head.Append(15)

	head = head.Prepend(10)
	head.Append(18)
	fmt.Println("First ListNode", head)
	head.Append(19)

	head = head.Prepend(0)
	fmt.Println("First ListNode", head)
	head.Traverse()

	fmt.Println("===============")

	head = head.Delete(0)
	head.Insert(3, 1, 2, 3)
	head.Traverse()
	head = head.Delete(7)
	head.Update(6, 100)
	head.Traverse()
	fmt.Println(head.Length())
	index, head2 := head.Search(3)
	fmt.Printf("Индекс => %d Элемент => %d\n", index, head2)
	head3 := head.Get(7)
	fmt.Println(head3)

}
