// slices2
// Make me compile!

package main

import "fmt"

func main() {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
	names_len := len(names)
	lastTwoNames := names[(names_len-2):(names_len)] // after figuring out the answer, try with other low/high bounds
	fmt.Println(lastTwoNames)
}
