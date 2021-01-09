package main

import "fmt"

func main() {
	s := "hello, 世界"
	buff := []rune("hello, 世界")

	fmt.Println(s[0])    // outputs 104
	fmt.Println(buff[0]) // outputs 104

	//s[0] = 'H'				// cannot assign to s[0]
	buff[0] = 'H'
	s = string(buff)

	fmt.Println(s[0]) // outputs 72
	fmt.Println(s)    // outputs Hello, 世界
}
