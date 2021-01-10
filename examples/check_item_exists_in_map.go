package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m) // map[James:32 Miss Moneypenny:27]

	fmt.Println(m["James"])    // 32
	fmt.Println(m["Barnabas"]) // 0

	v, ok := m["Barnabas"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // false

	if v, ok := m["Miss Moneypenny"]; ok { // ok will be false if the key doesn't exist in the map
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}
