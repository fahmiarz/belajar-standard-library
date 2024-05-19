package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(name, "Eko"))
	fmt.Println(slices.Contains(values, 100))
}
