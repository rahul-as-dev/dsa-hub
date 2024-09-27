package data_structures

import (
	"fmt"
	"math/rand/v2"
)

func Main6() {
	var s []int
	s = stackExample(s)
	s1 := make([]int, 0, 10)
	s1 = stackExample(s1)
	fmt.Println(s1)
}

// LIFO structure. Time complexity is O(1) for peek, pop, and push(amortized)
func stackExample(s []int) []int {
	//Push
	for range 5 {
		s = append(s, rand.IntN(100))
	}
	fmt.Println(s)
	//Peek
	peek := s[len(s)-1]
	fmt.Println(peek)
	//Polling
	s = s[:len(s)-1]
	fmt.Println(s)
	s = append(s, rand.IntN(100))
	fmt.Println(s)
	return s
}
