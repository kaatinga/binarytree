package main

type Tree struct {
	Item
}

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
	}

	if index > item.index {
		if item.right == nil {
			item.right = &Item{index: index, value: value}
		} else {
			item.right.Insert(index, value)
		}
	} else {
		if item.left == nil {
			item.left = &Item{index: index, value: value}
		} else {
			item.left.Insert(index, value)
		}
	}
}

// Find returns value of the index.
func (item *Item) Find(index byte) string {
	if item.index == index {
		return item.value
	}

	if index > item.index {
		return item.right.Find(index)
	}

	return item.left.Find(index)
}