package main

import (
	"fmt"
)

func main() {
	// var name = "Luan"
	// var version float32 = 1.1
	// age := 29
	// fmt.Println("Hello, Mr. ", name)
	// fmt.Println("This program is in version ", version)
	// fmt.Println("Your age is ", age)

	// fmt.Println("Type of variable 'name' is ", reflect.TypeOf(name))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido foi ", comando)
}
