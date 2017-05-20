package main

import "fmt"

func main() {
	subslice := memoryLeak()
	fmt.Println(subslice, " length: ", "remaining underlying array:", subslice[:cap(subslice)])
	noLeak := noMemoryleak()
	fmt.Println(noLeak, " length: ", "remaining underlying array:", noLeak[:cap(noLeak)])
}

func memoryLeak()  []int {
	s := []int{1,2,3,4,5,6,8,9,10}

	// Do some work

	return s[1:4]
}

func noMemoryleak() []int {
	s := []int{1,2,3,4,5,6,8,9,10}
	sub := make([]int, 3)
	copy(sub, s[1:4])

	return sub
}
