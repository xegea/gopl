// MissingInteger
// Converting array to map improves performance
// Score 66%: 	https://app.codility.com/demo/results/demoQGKZA2-7ZE/
// Score 100%: 	https://app.codility.com/demo/results/demo3U7KP4-MTK/
package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test2 := []int{1, 2, 4}
	test3 := []int{1, 2, 3}
	test4 := []int{-1, -3}
	test5 := []int{2}

	fmt.Println(test1, "-->", solution(test1))
	fmt.Println(test2, "-->", solution(test2))
	fmt.Println(test3, "-->", solution(test3))
	fmt.Println(test4, "-->", solution(test4))
	fmt.Println(test5, "-->", solution(test5))
}

func solution(A []int) int {
	m1 := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m1[A[i]] = A[i]
	}

	i := 1
	for ; i <= len(A); i++ {
		if _, ok := m1[i]; !ok {
			return i
		}
	}
	return i
}
