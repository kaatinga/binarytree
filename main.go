package main

import "fmt"

func main() {

	tree := Tree{}

	tree.Insert(128, "test")
	fmt.Println(tree.Find(128))
}
