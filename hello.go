package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Luan"
	var version float32 = 1.1
	age := 29
	fmt.Println("Hello, Mr. ", name)
	fmt.Println("This program is in version ", version)
	fmt.Println("Your age is ", age)

	fmt.Println("Type of variable 'name' is ", reflect.TypeOf(name))
}
