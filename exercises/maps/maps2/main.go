// maps2
// Make me compile!
//
package main

import "fmt"

func main() {
	m := map[string]int{
		"John" : 20, 
		"Ana" : 123,
	}
	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
