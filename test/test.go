package main

import (
	"slices"
)

func main() {
	test := []string{"1","2"}

	found := slices.Index(test, "3")

	println(found)
}
