package main

import (
	"fmt"
	sliceop "fops9311/go-sliceop"
)

func main() {
	s := []string{"123", "123", "321", "1515", "123"}

	filter := func(s string) bool {
		return s == "123"
	}
	sliceop.Filter(s, filter)

	fmt.Println(s) // [123 123 123]
}
