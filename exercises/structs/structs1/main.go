// structs1
// Make me compile!
//
package main

import "fmt"

type Person struct {
	name  string 
	age int32	
}

func main() {
	person := Person{
		name: "An",
		age: 12,
	}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
