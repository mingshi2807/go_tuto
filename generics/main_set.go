package main

import "fmt"

func main() {
	intSet := NewSet(1, 2, 3)
	fmt.Printf("Has 2 %v\n", intSet.Has(2))
	fmt.Printf("Has 5 %v\n", intSet.Has(5))

	strSet := NewSet("a", "b")

	fmt.Printf("Has a %v\n", strSet.Has("a"))
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}
