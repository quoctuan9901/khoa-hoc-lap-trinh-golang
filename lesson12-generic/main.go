package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Content     T
	Description T
}

func PrintValue[T any](value T) {
	fmt.Println(value)
}

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func IsNotEqual[T comparable](a, b T) bool {
	return a != b
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](a, b T) T {
	return a + b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func MaxLengthString(a, b string) string {
	if len(a) > len(b) {
		return a
	}

	return b
}

func main() {
	// PrintValue("Tuan")
	// PrintValue(100)
	// PrintValue(false)

	// result01 := IsEqual(5, 5)
	// PrintValue(result01)

	// result02 := IsEqual("Tuan", "Tuan")
	// PrintValue(result02)

	// result03 := IsEqual(3.5, 9.1)
	// PrintValue(result03)

	// result04 := IsNotEqual(3.5, 9.1)
	// PrintValue(result04)

	// PrintValue(Max(10, 7))
	// PrintValue(Max(5.5, 9.2))
	// PrintValue(Max("Tung", "Vy")) // ASCII

	// PrintValue(MaxLengthString("Tung", "Vy"))

	// stringBox := Box[string]{Content: "Hoc Golang Generic", Description: "Mo ta String Box"}
	// intBox := Box[int]{Content: 99, Description: 99}
	// PrintValue(stringBox.Content)
	// PrintValue(intBox.Content)

	// PrintValue(stringBox.Description)
	// PrintValue(intBox.Description)

	PrintValue(Sum(5.5, 10))
	PrintValue(Sum(9, 10))
	PrintValue(Sum(3.2, 6.4))
}
