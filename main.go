package main

import "fmt"

func main() {

	tree := Tree{}

	tree.Insert(128, "test128")
	fmt.Println(tree.Find(128))

	tree.Insert(127, "test127")
	fmt.Println(tree.Find(127))

	tree.Insert(126, "test126")
	fmt.Println(tree.Find(126))
}
