package main

import "math/rand"

func main() {

	tree := Tree{}

	tree.testValue(128)

	var number32, i int32
	for i < 600 {
		number32 = rand.Int31n(255)
		tree.testValue(byte(number32))
		i++
	}
}
