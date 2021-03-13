// Echo1 prints is command-line arguments
package main

import (
	"fmt"
	"strings"
	"time"
)

func echo1(arguments []string) float64 {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(arguments); i++ {
		s += sep + arguments[i]
		sep = " "
	}
	return time.Since(start).Seconds()
}

func echo2(arguments []string) float64 {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range arguments[1:] {
		s += sep + arg
		sep = " "
	}
	return time.Since(start).Seconds()
}

func echo3(arguments []string) float64 {
	start := time.Now()
	strings.Join(arguments[1:], " ")
	return time.Since(start).Seconds()
}

func main() {
	arguments := []string{}
	for i := 0; i < 40000; i++ {
		arguments = append(arguments, fmt.Sprint("param", i))
	}
	fmt.Println("Elapsed Times")
	fmt.Println("echo1 time: ", echo1(arguments))
	fmt.Println("echo2 time: ", echo2(arguments))
	fmt.Println("echo3 time: ", echo3(arguments))
}
