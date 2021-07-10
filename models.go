package main

import "fmt"

type Tree struct {
	*Item
}

func (tree *Tree) Insert(index byte, value string) {
	if tree.Item == nil {

		// the first item is added
		tree.Item = &Item{index: index, value: value}
		fmt.Println("root index was created:", index)
		fmt.Printf("%#v\n",tree.Item)
	} else {
		tree.Item.Insert(index, value)
	}
}

func (tree *Tree) testValue(index byte, value string) {
	tree.Insert(index, value)
	fmt.Println("returned value:", tree.Find(index))
	fmt.Println(counter)
	counter = 0
}

var counter uint

type Item struct {
	index byte
	value string
	left  *Item
	right *Item
}

// Insert inserts new item into the btree.
func (item *Item) Insert(index byte, value string) {

	if index == item.index {
		item.index = index
		item.value = value
		fmt.Println("current index was updated:", index)
	}

	if index > item.index {
		if item.right == nil {
			item.right = &Item{index: index, value: value}
			fmt.Println("new index was created:", index)
		} else {
			item.right.Insert(index, value)
		}
	} else {
		if item.left == nil {
			item.left = &Item{index: index, value: value}
			fmt.Println("new index was created:", index)
		} else {
			item.left.Insert(index, value)
		}
	}
}

// Find returns value of the index.
func (item *Item) Find(index byte) string {

	counter++

	if item == nil {
		return "not found"
	}

	if item.index == index {
		return item.value
	}

	if index > item.index {
		return item.right.Find(index)
	}

	return item.left.Find(index)
}
