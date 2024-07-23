package main

import (
	"fmt"
	"math"
)

func main() {
	PrintBalance(NZD(250))
}

type Currency interface {
	~int | ~int64
	ISO4127Code() string
	Decimal() int
}

func PrintBalance[T Currency](b T) {
	balance := float64(b) / math.Pow10(b.Decimal())
	fmt.Printf("%.*f %s\n", b.Decimal(), balance, b.ISO4127Code())
}

type NZD int64

func (b NZD) ISO4127Code() string {
	return "NZD"
}

func (b NZD) Decimal() int {
	return 2
}
